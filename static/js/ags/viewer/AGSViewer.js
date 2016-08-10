

Ext.define('ags.viewer.AGSViewer' ,{
    extend: 'Ext.panel.Panel',
	requires: [
		//"Ags.abbrev.MetaStore"
	],

	initComponent: function(){


		 Ext.apply(this, {
		 	renderTo: "widget_div",
			title : 'File ',
			height: HEIGHT,
  			layout: "border",

  			tbar: [
  				//"->",
  				{text: "Load", scope: this, handler: this.fetch}
  			],


			items: [
				this.get_tab_panel()
			]
		});
		this.callParent();
	},

	get_tab_panel: function(){
		if(!this._tabpanel){
			this._tabpanel = Ext.create("Ext.tab.Panel", {
				flex: 1,
				region: "center"
			});
		}
		return this._tabpanel
	},

	fetch: function(){
		Ext.MessageBox.wait('Loading...');
		Ext.Ajax.request({
			scope: this,
			url: '/ags/4/parse?example=03-total_station_point_4_0.ags',
			params: {

			},

			success: function(response){

				var data = Ext.decode(response.responseText);
				//console.log(data);
				var groups = data.document.groups;

				for(var i = 0; i < groups.length; i++){


					var col_defs = [];
					var model_fields = [];
					//var columns = [];

					var tab = Ext.create("ags.viewer.GroupView", {
                    							g_title: grp.group_code,
                    							g_tooltip: grp.description,
                    							g_itemId: grp.group_code,
                    							g_columns: col_defs,
                    							g_store: sto

                    });
					this.get_tab_panel().add(tab)
                    continue;

					var grp = groups[i];


					//var row_count = -1;
					// Create headings
					var headings_len =  grp.headings.length
					for(var c = 0; c < headings_len; c++){

						var h = grp.headings[c];
						// add field def to model
						model_fields.push( {name: h.head_code, type: "string"} );

						// col def for grid, also hide data in `geo_data`
						var col = {header: h.head_code, dataIndex: h.head_code,
									sortable: true, menuDisabled: true,
									head_code: h.head_code
									};
						col_defs.push(col);
					}
					//console.log(grp.group_code,  h.head_code + ".value");

					var model = this.make_model(model_fields);
					var sto = Ext.create("Ext.data.Store", {model: model});
					for(var di=0; di < grp.data.length; di++){
						var rd = grp.data[di];
						var	rec = {};
						for(var cd = 0; cd < headings_len; cd++){
							var hhc = grp.headings[cd].head_code;
							rec[hhc] = rd[hhc].value;
						}
						sto.add(rec);
					}


					var tab = Ext.create("ags.viewer.GroupView", {
							g_title: grp.group_code,
							g_tooltip: grp.description,
							g_itemId: grp.group_code,
							g_columns: col_defs,
							g_store: sto

					});
					this.get_tab_panel().add(tab)

				}

			}
		});
		Ext.MessageBox.hide();
	},

	make_model: function(fields){

		return Ext.define('Ags.dymamic.MODEL' + Ext.id(), {
			extend: 'Ext.data.Model',
			fields: fields,
			proxy: {
				type: 'memory',
				reader: {
					type: 'json',
					totalProperty: 'tc',
					root: 'foobar'
				}
			}
		});
	}



});