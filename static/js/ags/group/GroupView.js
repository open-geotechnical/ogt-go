
Ext.define('Ags.group.GroupView' ,{
    extend: 'Ext.panel.Panel',
	requires: [
		//"Ags.abbrev.MetaStore"
	],

	initComponent: function(){


		 Ext.apply(this, {
			title : 'Group ',
			height: HEIGHT,
  			layout: "border",
			items: [
				Ext.create("Ags.group.HeadingsGrid", {flex: 1, region: "center"})
			]
		});
		this.callParent();
	}



});