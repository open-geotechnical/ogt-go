
Ext.define('ags.group.GroupView' ,{
    extend: 'Ext.panel.Panel',
	requires: [
		//"Ags.abbrev.MetaStore"
	],

	initComponent: function(){


		 Ext.apply(this, {
			title : 'Group: ',
			height: HEIGHT,
  			layout: "border",
			items: [
			    {xtype: "tabpanel",
			        items: [
				        Ext.create("ags.group.HeadingsGrid", {
				        title: "Headings", flex: 1, region: "center"})
				    ]
				}
			]
		});
		this.callParent();
	}



});