import sys
from unagent_lint.runner import run

if len(sys.argv) < 2:
    print("Path required")
    sys.exit(3)

report = run(sys.argv[1])
if report["failures"]:
    sys.exit(1)
if report["warnings"]:
    sys.exit(2)
