

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
			this._tabpanel = Ext.create("Ext.tab.Panel", {flex: 1, region: "center",

				deaditems: [
					{title: "Source"}

				]
			});


		}
		return this._tabpanel

	},

	fetch: function(){

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
					/*
					var rows = []
					for(var r = 0; r < row_count; r++){

						var row = {};
						for(var c = 0; c < col_defs.length; c++){
							var h = col_defs[c];

							row[h.head_code] = h.geo_data[r].value;
							//console.log("=", h, h.geo_data[r], row)
						}
						rows.push(row);
						console.log("r=", r, row);
					}
					*/
					//console.log(rows);
					//var val_rows = [];
					//var data = grp.data;
					//for(var i=0; i < data; i++){
					//	var rd = data[i];
					//	row
					//	for(var prop in rd){
					//		if (rd.hasOwnProperty(prop)) {

					//		}
					//	}
					//}

					var model = this.make_model(model_fields);
					var sto = Ext.create("Ext.data.Store", {model: model});
					for(var di=0; di < grp.data.length; di++){
						var rd = grp.data[di];
						var	rec = {};
						for(var cd = 0; cd < headings_len; cd++){
							var hhc = grp.headings[cd].head_code;
							rec[hhc] = rd[hhc].value;
						}
						//console.log(rec)
						sto.add(rec);
					}


					//;
					//sto.loadData(grp.data);

					var tab = Ext.create("ags.viewer.GroupGrid", {
							title: grp.group_code,
							itemId: grp.group_code,
							columns: col_defs,
							store: sto

					});
					this.get_tab_panel().add(tab)


				}

			}
		});
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