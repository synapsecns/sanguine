import time
import subprocess
import sys
from datetime import datetime

callStatement = sys.argv[1]

while 1==1:
	subprocess.call(f"{callStatement}", shell=True)
	print (f"{datetime.now()} - Restarting...")
	time.sleep(1)