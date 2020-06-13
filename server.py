
import SimpleHTTPServer
import SocketServer
import psutil
import json
import subprocess

class MyHttpRequestHandler(SimpleHTTPServer.SimpleHTTPRequestHandler):
    def do_GET(self):
        if self.path == '/':
            self.send_response(200)
            self.end_headers()

            #cpu_util = psutil.cpu_percent()
            memory_util = psutil.virtual_memory()[2]
            #connect = psutil.net_connections(kind='udp')
            data_set = { 
                         "Cpu_Utilization" : psutil.cpu_percent() ,

                         "Memory_Utilization" : memory_util ,
                         "Network_Utilization_out" : psutil.net_io_counters(pernic=False,nowrap=False)[0],
                         #"Network_Utilization_in" : psutil.net_io_counters(pernic=False,nowrap=True)[1]
                         "Connections "   : subprocess.check_output('netstat -ano | grep ESTABLISHED | wc -l', shell=True)
                         
                         }
            json_dump = json.dumps(data_set) #string of json object

            #json_object = json.loads(json_dump)
            
            #self.wfile.write(type(json_dump))
            #self.wfile.write(type(data_set))
            self.wfile.write(json_dump)
            #self.wfile.write((data_set[]))
            #self.wfile.write(type(memory_util))
            #elf.wfile.write(type(connect))
            #self.wfile.write(type(data_set))
            
            
            
            

            

        else:

            self.send_error(404, notfound)
# Create an object of the above class

handler_object = MyHttpRequestHandler

PORT = 9999 #change code accordingly
my_server = SocketServer.TCPServer(("", PORT), handler_object)

# Start the server
my_server.serve_forever()