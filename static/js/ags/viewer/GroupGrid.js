Ext.define('ags.viewer.GroupGrid' ,{
    extend: 'Ext.grid.Panel',
	requires: [
		//"Ags.abbrev.AbbrevsStore"
	],

	initComponent: function(){
		 Ext.apply(this, {
			deadtitle : 'GROUP',
			//store: Ext.getStore("abbrevs"),
			height: HEIGHT,

			deadcolumns: [
				{header: 'Heading',  dataIndex: 'head_code',  flex: 1, menuDisabled: true, sortable: true,
					renderer: R.bold
				},
				{header: 'Description', dataIndex: 'description', flex: 1, menuDisabled: true, sortable: true},
				{header: 'Group', dataIndex: 'group', flex: 1, menuDisabled: true, sortable: true}
			],

			deaddockedItems: [{
                    xtype: 'pagingtoolbar',
                    //store: Ext.getStore("abbrevs"),
                    dock: 'bottom',
                    displayInfo: true
			}],

			listeners: {
				select: function(obj, rec, opts){

				}
			}
		});
		this.callParent();
	}



});