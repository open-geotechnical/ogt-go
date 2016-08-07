
Ext.define('Ags.group.HeadingsGrid' ,{
    extend: 'Ext.grid.Panel',
	requires: [
		//"Ags.abbrev.AbbrevStore"
	],

	initComponent: function(){
		 Ext.apply(this, {
			title : 'Groups Index',
			store: Ext.getStore("groups"),
			height: HEIGHT,

			columns: [
				{header: 'Heading', dataIndex: 'code', flex: 1, menuDisabled: true, sortable: true, renderer: R.bold},
				{header: 'Description', dataIndex: 'description', flex: 3, menuDisabled: true, sortable: true},
				{header: 'Class',  dataIndex: 'class',  flex: 2, menuDisabled: true, sortable: true}

			],

			dockedItems: [{
                    xtype: 'pagingtoolbar',
                    store: Ext.getStore("headings"),
                    dock: 'bottom',
                    displayInfo: true
			}],

			listeners: {
				DEADDDDDselect: function(obj, rec, opts){
					console.log("yes", rec, rec.get("heading"));
					var sto = Ext.getStore("abbrev_items");
					//console.log("detch", rec);
					//var sto = Ext.getStore("abbrev_items");

					//var proxy = Ext.getStore("abbrev_items").getProxy()
					//console.log("proxy", proxy);
					sto.getProxy().url = "/ags/4/abbrev/" + rec.get("heading");
					sto.load()
				}
			}
		});
		this.callParent();
	}



});