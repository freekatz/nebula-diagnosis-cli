## Logger bugs
INFO[0001] node0: physical                              
INFO[0002] node0: physical                              
INFO[0003] node0: physical                              
INFO[0003] node1: physical // should be `INFO[0001] node1: physical` because the node1 run with a new goroutine               
INFO[0004] node0: physical                              
INFO[0005] node0: physical                              
INFO[0006] node1: physical  // should be `INFO[0002] node1: physical` because the node1 run with a new goroutine                           
INFO[0006] node0: physical                              
INFO[0007] node0: physical                              
INFO[0008] node0: physical                              
INFO[0009] node1: physical // should be `INFO[0003] node1: physical` because the node1 run with a new goroutine                             
INFO[0009] node0: physical                              
INFO[0010] node0: physical                              
INFO[0011] node0: physical                              
INFO[0012] node0: physical                              
INFO[0013] node0: physical                              
INFO[0014] node0: physical                              
INFO[0015] node0: physical                              
INFO[0016] node0: physical                              
INFO[0017] node0: physical


