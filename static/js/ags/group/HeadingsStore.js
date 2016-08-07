
Ext.define('Ags.group.HeadingsStore', {
	extend: 'Ext.data.Store',
	requires: [
       //'Ags.model.Abbrev'
    ],
	constructor: function(){
		Ext.apply(this, {
			model: 'Ags.model.Heading',
			storeId: "headings",
			sssorters: [ {
				property: 'dated',
				direction: 'DESC'
			}],
			deadsortInfo: {
				property: 'code',
				direction: 'DESC'
			},
			groupField: "group",
			pageSize: 1000,
			autoLoad: true,
			proxy: {
				type: 'ajax',
				//url: "/ags/4/groups.json",
				reader: {
					type: 'json',
					root: "groups",
					idProperty: 'code',
					sstotalProperty: 'code_count'
				}
			}
		});
		this.callParent();

	}
});
