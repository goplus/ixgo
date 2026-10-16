"""Serial comparison; writes local language_timings.json next to this file."""
import json
import platform
import shutil
import statistics
import subprocess
import sys
import time
from pathlib import Path

root = Path(__file__).resolve().parent
available = {"python": "python3", "ruby": "ruby", "lua": "lua", "luajit": "luajit"}
selected = sys.argv[1:] or ["python", "ruby", "lua", "luajit"]
commands = {name: [shutil.which(available[name])] for name in selected}
if any(not command[0] for command in commands.values()):
    raise RuntimeError("Selected interpreters must all be installed")
report = {"platform": platform.platform(), "input": 40, "rounds": 3, "method": "Serial process wall time including startup; no warmup; no JIT flags", "versions": {}, "results": {}}
if (root / "language_timings.json").exists():
    previous = json.loads((root / "language_timings.json").read_text())
    if previous.get("platform") == report["platform"]:
        report = previous
for language, command in commands.items():
    flag = "-v" if language in ("lua", "luajit") else "--version"
    report["versions"][language] = subprocess.check_output(command + [flag], stderr=subprocess.STDOUT, text=True).strip()
    suffix = {"python": "py", "ruby": "rb", "lua": "lua", "luajit": "lua"}[language]
    if language == "luajit":
        report["luajit_status"] = subprocess.check_output(command + ["-e", 'print(jit.version); print(jit.status())'], text=True).strip()
    # Check base cases and nontrivial small inputs before timing.
    for example in ("fibish", "fibish_closure"):
        script = root / f"{example}.{suffix}"
        for n, pair in ((0, (0, 0)), (1, (0, 1)), (2, (0, 2)), (3, (2, 1)), (6, (3, 10))):
            got = subprocess.check_output(command + [str(script), str(n)], text=True).strip()
            assert got == f"f({n})={pair[0]},{pair[1]}", got
        key = f"{language}/{example}"
        runs = []
        for round in range(3):
            start = time.perf_counter()
            got = subprocess.check_output(command + [str(script)], text=True).strip()
            elapsed = time.perf_counter() - start
            assert got == "f(40)=39088169,126491972", got
            runs.append(elapsed)
            print(f"{key} round {round+1}: {elapsed:.6f}s; {got}", flush=True)
        report["results"][key] = {"seconds": runs, "median": statistics.median(runs), "output": got}
        (root / "language_timings.json").write_text(json.dumps(report, indent=2) + "\n")
        print(f"{key} median: {statistics.median(runs):.6f}s", flush=True)
