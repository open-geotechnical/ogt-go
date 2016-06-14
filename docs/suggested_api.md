AGS and Pedro.bill brief
==================

There are a few parts to this project..
and a load of legacy which mashed up makes some sense..
cos it not weel explained a and bit of a dark art..

- The AGS File format which compirises of
- a HEADING
- which has an UNIT, eg Temp degrees C
- and a datatype? eg number to 3 decimals

And all the above need to be validated for errors,
data definition mistakes et all...

So here is Pete's plan

First we start at AGS.v4 which is latest and greatest

The "rip/scrape" is in ags-data-json
Which is WRONG should be the ags-spec-json

Units
===========
endpoint = http://ags.daffodil.uk.com/spec/4/units
or
http://aclontrol.gstl.co.uk/standards/ags/spec/4/units
data atmo is in BB/ags-data-json/4/units.json
Which is a scrape

We start with the UNITS of measure..
such as 
- temp C vs F
- gallons vs litres
- and non metric
- still used in the USA and alike

But main thing we want to do is to normalise the units in a few ways..
and the daffodil speciallity of adding formulates and functions..
to the units...

So each unit we really need to focus upon and understand from and engineer perspective..
and this would mean adding "meta data" and links and sone stuff.. eg 

{
  "description": "percentage", 
  "unit": "%"
}, 

Daffodil add stuff like
{
  "description": "percentage", 
  "unit": "%",
  "unicode" "%",
  links:
     - "standard": "https:://some.standard.body",
     - wikipedia: wikipedia..
  description:
    -- and off we go to description and even calculation
  }, 

So we need to get funky on this..
And Big Bill..
We actually knock it out in code in tests..
and make a whole "Unit" a type..

Eg 
Temp as type..
temp := Temp(str, CENTIGARE)
temp.InCentrigade()
temp.InKelvin()
temp.InF()
temp.ToSecyString()

Dont matter..
We NEED TO STICK to ags standard..
So sometimes we roll out our own..
and sometimes we use externals..


















