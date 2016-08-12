

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
  				{text: "Open Example", scope: this, handler: this.select_example},
  				"-",
  				{text: "Upload AGS", scope: this, handler: this.upload_dialog},
  				"->",
  				{text: "Clear", scope: this, handler: this.clear_all},

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

	clear_all: function(){
		var tabPanel = this.get_tab_panel();
		tabPanel.removeAll();
    },

	load_example: function(file_name){

		this.clear_all();

		Ext.MessageBox.wait('Loading...');
		//this.mask("")
		Ext.Ajax.request({
			scope: this,
			url: '/ajax/ags/4/parse',
			method: "GET",
			params: {
				example: file_name
			},

			success: function(response){

				var data = Ext.decode(response.responseText);
				var tabPanel = this.get_tab_panel();
				var groups = data.document.groups;

				for(var i = 0; i < groups.length; i++){


					var col_defs = [];
					var model_fields = [];
					//var columns = [];

					var tab = Ext.create("ags.viewer.GroupView", {});
                    tab.load_group(groups[i])
					tabPanel.add(tab)
					if(i == 0){
						tabPanel.setActiveTab(tab);
					}

				}
				Ext.MessageBox.hide();
			}

		});

	},

	select_example: function(){
		var d = Ext.create("ags.examples.ExamplesDialog", {});
		d.on("OPEN", function(fn){
			this.load_example(fn);
		}, this)
		d.load_show();

	},

	upload_dialog: function(){
    		var d = Ext.create("ags.viewer.UploadDialog", {});
    		d.on("OPEN", function(fn){
    			this.load_example(fn);
    		}, this)
    		d.show();
    }

});