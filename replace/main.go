package main

//removes unnecessary characters from an email address list separated by semicolons

import (
	"fmt"
	"strings"
)

//copy your email address into the myaddress function declaration
func main() {
	myaddress := "(agrosman@lmcm.com); amorey@leemunder.com; (awetzel@flputnam.com);  (andrew.wigren@fmr.com); andrew.wright@blackrock.com; (blandy@greenhousefunds.com);  (benjamin.piggott@fmr.com);  (wmogrodnick@wellington.com); bob.bertelson@fmr.com; boyd.jm@tbcam.com; cely@nicholsassetmgmt.com; charlie.hebard@fmr.com;  (clinton.condon@fmglobal.com);  <dfarb@highfieldscapital.com>;  (dneumann@grtcapital.com); dan.hagen@peregrinecapital.com; danoff-research@fmr.com;  (dwpalmer@wellington.com); dbisson@centurycap.com; drice@grtcapital.com; dsmith@nicholsassetmgmt.com; edward.davis2@fidelity.com; efischman@mfs.com;  <Jonathan_Evans@troweprice.com>; forrest.stclair@fmr.com;  (Jerry@Monasheecap.com); GJKomansky@ClearBridgeAdvisors.com;  (gleblanc@wellington.com); harmon.research@fmr.com; ira@brantpointfund.com; jacob_kann@troweprice.com; jim.ross@peregrinecapital.com; jnmordy@wellington.com;  (jbozoyan@manulifeam.com); joel.tillinghast@fmr.com;  (john_brellenthin@decade-llc.com);  (john.dowd@fmr.com); johnr.murray@pioneerinvestments.com; jonathan.kasen@fmr.com; jswalker@mfs.com; jwilladsen@leemunder.com; kebandtel@wellington.com; ken.winston@pioneerinvestments.com; kjackson@mfs.com; kthomson@mfpllc.com; leannemoore143@gmail.com; lionel.harris@fmr.com; lthorndike@centurycap.com; mgrossman@mfs.com; mnviviano@wellington.com; nathan.strik@fmr.com; neal.p.miller@fmr.com; nicole.esposito@fmr.com; nmchoumenkovitch@wellington.com; Nwang@Northruncapital.com <Nwang@Northruncapital.Com>; <PJCollier@ClearBridgeAdvisors.com>; phperelmuter@wellington.com; piskorowski.jj@tbcam.com;  (philmuscatello@gmail.com); rgillis@nicholsassetmgmt.com; rgranahan@granahan.com;  (romain@juliencapital.com); rsatterfield@leemunder.com; (rmcallister@mfs.com);  (ryan.oldham@fmr.com); (sandeep.gupta@pyramis.com);  (tbirdvergin@gmail.com); sgreen@nicholsassetmgmt.com;  <Riaz.research@fidelity.com>; Shadman.riaz@fmr.com; shaji.john@pioneerinvestments.com; (shawn_driscoll@troweprice.com); sonu.kalra@fmr.com; tellis@northruncapital.com; thomas.soviero@fmr.com;  <tom.bushey@sunderlandcapital.com>; tvingers@leemunder.com; wehbe.r@tbcam.com; will.danoff@fmr.com; william.taylor@pioneerinvestments.com"
	myaddress = strings.Replace(myaddress, "(", "", -1)
	myaddress = strings.Replace(myaddress, ")", "", -1)
	myaddress = strings.Replace(myaddress, "<", "", -1)
	myaddress = strings.Replace(myaddress, ">", "", -1)
	fmt.Println(myaddress)
}
