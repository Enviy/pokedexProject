#!/usr/local/bin/python

import sys, os

allowedRequests = ["root"]

if len(sys.argv) != 2:
    print("Invalid Number of Parameters, Need 1")
    sys.exit(1)

if sys.argv[1] not in allowedRequests:
    print("The passed in request is not in the list of supported requests: " + str(allowedRequests))
    sys.exit(1)

os.system("python ./canned-requests/" + sys.argv[1] + ".py")

