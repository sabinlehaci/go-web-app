import os
import requests
import sys

TOKEN= str(sys.argv[0])



print( "the token value is")
def trigger_workflow(Workflow_Name):

      headers = {
        "Accept": "application/vnd.github.v3+json",
        "Authorization": f"token {TOKEN}",
      }

      data = {
        "event_type": Workflow_Name,
        "client_payload": {
          'baselinetag': pl_Baseline_Number,
          'revision_number': pl_Baseline_Revision
        }
      }

      responsevalue=requests.post(f"https://api.github.com/repos/hashicorp/immutable-infrastructure/dispatches",json=data,headers=headers)
      print("The response message is ",responsevalue.content)

trigger_workflow(Workflow_Name)
