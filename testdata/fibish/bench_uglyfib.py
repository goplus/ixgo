"""Serial uglyfib comparison; writes local uglyfib_timings.json."""
import json
import platform
import shutil
import statistics
import subprocess
import sys
import time
from pathlib import Path

root = Path(__file__).resolve().parent
report = {"platform": platform.platform(), "input": 40, "rounds": 3, "results": {}}
if (root / "uglyfib_timings.json").exists():
    previous = json.loads((root / "uglyfib_timings.json").read_text())
    if previous.get("platform") == report["platform"]:
        report = previous
report["method"] = "Serial wall time including startup; shared mutable tables/lists/arrays model scalar pointers and local arrays; no JIT flags"
interpreters = {"lua": ("lua", "lua"), "luajit": ("luajit", "lua"), "python": ("python3", "py"), "ruby": ("ruby", "rb")}
for name in sys.argv[1:] or interpreters:
    executable, suffix = interpreters[name]
    binary = shutil.which(executable)
    if binary is None:
        raise RuntimeError(f"Interpreter not installed: {executable}")
    flag = "-v" if name in ("lua", "luajit") else "--version"
    version = subprocess.check_output([binary, flag], stderr=subprocess.STDOUT, text=True).strip()
    script = root / ("uglyfib." + suffix)
    for n, expected in ((0, 0), (1, 1), (2, 1), (3, 2), (6, 8), (10, 55), (20, 6765)):
        got = subprocess.check_output([binary, str(script), str(n)], text=True).strip()
        assert got == f"Fib({n})={expected}", got
    runs = []
    for i in range(3):
        start = time.perf_counter()
        got = subprocess.check_output([binary, str(script)], text=True).strip()
        elapsed = time.perf_counter() - start
        assert got == "Fib(40)=102334155", got
        runs.append(elapsed)
        print(f"{name} round {i+1}: {elapsed:.6f}s; {got}", flush=True)
    report["results"][name] = {"version": version, "seconds": runs, "median": statistics.median(runs), "output": got}
    (root / "uglyfib_timings.json").write_text(json.dumps(report, indent=2) + "\n")
    print(f"{name} median: {statistics.median(runs):.6f}s", flush=True)
