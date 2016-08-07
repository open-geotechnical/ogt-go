
var VP = 0;

var HEIGHT = window.innerHeight - 110;

/*
Ext.Loader.setConfig({
	enabled: true,
	paths: {
		'Ags': '/js/ags'
		//'Ext.ux': "/js/ux"
	}
});
*/


Ext.create("Ags.abbrev.AbbrevStore", {});
Ext.create("Ags.abbrev.AbbrevItemsStore", {});

R = {}
R.bold = function(val, meta, xxx){
	return "<b>" + val + "</b>";
}