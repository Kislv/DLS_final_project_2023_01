import sys
from concurrent import futures
import grpc

import countPores 
sys.path.append( '../../GRPC/CountPores/ML' )
import countPores_pb2
import countPores_pb2_grpc



class CountPoresServicer(countPores_pb2_grpc.CountPoresServicer):

  def Count(self, request, context):
    area, diameter = countPores.count_pores(request.extenstion, (request.width, request.height), request.image, request.density)
    return countPores_pb2.CountPoresResponse(area = float(area), diameter = float(diameter))
  
def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  countPores_pb2_grpc.add_CountPoresServicer_to_server(
      CountPoresServicer(), server)
  server.add_insecure_port('[::]:40042')
  server.start()
  print("Pores segmentation mcs start!")
  server.wait_for_termination()

serve()
