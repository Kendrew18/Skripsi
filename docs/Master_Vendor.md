# How TO Use API Master Vendor
__________
##  Input Master Vendor

Link: kostsoda.onthewifi.com:3333/MV/input-master-vendor

Method: POST

Controllers:

    nama_vendor := c.FormValue("nama_vendor")
    pekerjaan_vendor := c.FormValue("pekerjaan_vendor")

##  Read Master Vendor

Link: kostsoda.onthewifi.com:3333/MV/read-master-vendor

Method: POST

Controllers: -

##  Edit Master Vendor

Link: kostsoda.onthewifi.com:3333/MV/edit-master-vendor

Method: PUT

Controllers:

    id_vendor := c.FormValue("id_vendor")
	nama_vendor := c.FormValue("nama_vendor")
	pekerjaan_vendor := c.FormValue("pekerjaan_vendor")

##  Delete Master Vendor

Link: kostsoda.onthewifi.com:3333/MV/delete-master-vendor

Method: Delete

Controllers:

    id_vendor := c.FormValue("id_vendor")