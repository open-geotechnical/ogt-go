ogt-ags-go
===================

The AGS file format is a data exchange format
used extensively used in GeoTechnical engineering, and almost a defacto standard.

This is R&D with ags4 data format
in golang, so please have fun.

Disclaimer.. This is experimenting with an online server and rapid ags ref server upon IOT..
please have fun.. could it be usable ??

[![GoDoc](https://godoc.org/github.com/open-geotechnical/ogt-ags-go?status.svg)](https://godoc.org/github.com/open-geotechnical/ogt-ags-go)


Dev Guide
==============

So far, this is an application is that contains the ags4 data spec.

The language and server process is in golang, and upon startup it reads
from the ags-datadict, the whole lot into memory...
and thaen the webserver waiting for requests..

