# How TO Use API Jadwal
__________
##  Input Durasi Pekerjaan

Link: kostsoda.onthewifi.com:3333/PJDL/input-durasi-task

Method: PUT

Controllers:

    id_penjadwalan := c.FormValue("id_penjadwalan")
	waktu_optimis := c.FormValue("waktu_optimis")
	waktu_pesimis := c.FormValue("waktu_pesimis")
	waktu_realistis := c.FormValue("waktu_realistis")

## Read Task

Link: kostsoda.onthewifi.com:3333/PJDL/read-task

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

## Read Depedentcies

Link: kostsoda.onthewifi.com:3333/PJDL/read-dep

Method: GET

Controllers:
    
     id_proyek := c.FormValue("id_proyek")
     id_penjadwalan := c.FormValue("id_penjadwalan")

## Input Depedentcies

Link: kostsoda.onthewifi.com:3333/PJDL/input-depedentcies

Method: PUT

Controllers:

    depedentcies := c.FormValue("depedentcies")
    id_jadwal := c.FormValue("id_jadwal")

NB: Contoh Depedentcies seperti -> |Task F|

##  Generate Jadwal

Link: kostsoda.onthewifi.com:3333/PJDL/generate-jadwal

Method: PUT

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Read Jadwal

Link: kostsoda.onthewifi.com:3333/PJDL/read-jadwal

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Edit Durasi dan Tanggal

Link: kostsoda.onthewifi.com:3333/PJDL/edit-dur-tgl

Method: PUT

Controllers:

    id_penjadwalan := c.FormValue("id_penjadwalan")
    tanggal_pekerjaan_mulai := c.FormValue("tanggal_pekerjaan_mulai")
    durasi := c.FormValue("durasi")

##  See Calender

Link: kostsoda.onthewifi.com:3333/PJDL/see-calender-all

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")
    status_user := c.FormValue("status_user")
