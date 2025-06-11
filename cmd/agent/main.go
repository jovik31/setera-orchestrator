package main

/*TODO: Implement the agent main function
[ ]: Agents needs to run a daemonset to control the following:
	[ ]: Nodestore creation
		[ ] Node labels and annotations
        [ ] Node status, tolerations and taints
	[ ]: Nodestore update
        [ ] Check if there is an update to the node annotations, labels and tolerations
		[ ] If full add node full for scheduling

[ ]: Run a unix domain socket server to communicate with the CNI plugin for the following:
	[ ]: Serve requests for pod information for tenant info:
	 	[ ]: Check if tenant exists
		[ ]: Check if tenant has 
*/
func main(){



	
	/*

	Agent runs as a daemonset
	Initializes and creates the nodestore corresponding to the node where it is running
	Starts running the the agent that watches all Tenants and a single Nodestore.
	It updates the nodestore with all node network information. 
	It updates the tenant with its network information 
	
	*/

	/* Create nodestore */

	

}