
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
			autoLoad: false,
			proxy: {
				type: 'ajax',
				url: "/ags/4/null",
				reader: {
					type: 'json',
					root: "group.headings",
					idProperty: 'head_code',
					sstotalProperty: 'code_count'
				}
			}
		});
		this.callParent();

	}
});
