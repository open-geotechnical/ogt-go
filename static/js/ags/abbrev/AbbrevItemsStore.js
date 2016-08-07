
Ext.define('Ags.abbrev.AbbrevItemsStore', {
	extend: 'Ext.data.Store',
	requires: [
       //'Ags.model.AbbrevItem'
    ],
	constructor: function(){
		Ext.apply(this, {
			model: 'Ags.model.AbbrevItem',
			storeId: "abbrev_items",
			pageSize: 1000,
			autoLoad: false,
			proxy: {
				type: 'ajax',
				reader: {
					type: 'json',
					root: "abbreviation.items",
					idProperty: 'item',
					sstotalProperty: 'code_count'
				}
			}
		});
		this.callParent();
	},

	deadfetch: function(rec){
		console.log("detch", rec);
		//var sto = Ext.getStore("abbrev_items");

		//var proxy = Ext.getStore("abbrev_items").getProxy()
		//console.log("proxy", proxy);
		this.getProxy().url = "/ags/4/abbrev/" + rec.get("heading");
		this.load()
	}
});