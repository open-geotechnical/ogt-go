
Ext.define('Ags.model.Abbrev', {
    extend: 'Ext.data.Model',

    fields: [
		'head_code',
		'description',
		'group'
	]
});

Ext.define('Ags.model.AbbrevItem', {
    extend: 'Ext.data.Model',

    fields: [
		'item',
		'description',
		"date_added",
		"added_by",
		"status"
	]
});

Ext.define('Ags.model.Group', {
    extend: 'Ext.data.Model',

    fields: [
		'group_code',
		'description',
		'class'
	]
});
Ext.define('Ags.model.Heading', {
    extend: 'Ext.data.Model',

    fields: [
		'head_code',
		'description',
		'date_type',
		'unit',
		'example',
		'status',
		'sort',
		'rev_date',

	]
});


Ext.create("Ags.abbrev.AbbrevsStore", {});
Ext.create("Ags.abbrev.AbbrevItemsStore", {});

Ext.create("Ags.group.GroupsStore", {});
Ext.create("Ags.group.HeadingsStore", {});
