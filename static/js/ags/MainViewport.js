Ext.define('Ags.MainViewport', {
	extend: "Ext.tab.Panel",
	renderTo: "widget_div",
	width: "100%",
	//height: window.innerHeight - 80,
	plain: true,
	dtitle: "AGS 4",
	border: 0,
    layout: 'fit',
    activeTab: 0,
    items: [

			{xtype: "panel", title: "Abbreviations", layout: "column", border: 0, plain: true,
				items: [
					Ext.create("Ags.abbrev.AbbrevsGrid", {columnWidth: 0.5, id: "AbbrevsGrid"}),
					Ext.create("Ags.abbrev.AbbrevView", {columnWidth: 0.5, id: "AbbrevView"})
				]
			},
			{flex: 10, 
					id: "iframe_xid",
			        xtype : "component",
					title: "Docuementation",
			        autoEl : {
			            tag : "iframe", 
			            src : "https://agsngx.readthedocs.org/en/latest/"
			        }
		    }
        ]
});