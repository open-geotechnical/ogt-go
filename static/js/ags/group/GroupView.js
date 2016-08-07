
Ext.define('Ags.group.GroupView' ,{
    extend: 'Ext.panel.Panel',
	requires: [
		//"Ags.abbrev.MetaStore"
	],

	initComponent: function(){


		 Ext.apply(this, {
			title : 'Group ',
			height: HEIGHT,
  			layout: "vbox",
			items: [
				//Ext.create("Ags.group.HeadingsGrid", {})
			]
		});
		this.callParent();
	}



});