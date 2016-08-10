
Ext.define('ags.viewer.GroupView' ,{
    extend: 'Ext.panel.Panel',
	requires: [
		//"Ags.abbrev.AbbrevsStore"
	],

	initComponent: function(){
		 Ext.apply(this, {
			//deadheight: HEIGHT,
			layout: "border",
			title: this.g_title,
			tooltip: this.g_description,
			items: [
				{xtype: "displayfield", name: "description", value: "-dsadsadsa--", region:"north", sflex: 1},
				this.get_grid(this.g_columns, this.g_store)
			]

		});
		this.callParent();
	},

	get_grid: function(columns, store){

		if(!this._grid){
			this._grid = Ext.create("ags.viewer.GroupGrid", {
			dddheight: HEIGHT - 110,
				autoHeight: true,
				region:"center",
				ssflex: 5,
				columns: columns,
				store: store
			});
		}
		return this._grid;
	}

});