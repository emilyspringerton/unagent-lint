from unagent_lint.rules import RULES


def run(path):
    with open(path, encoding="utf-8") as handle:
        content = handle.read()

    violations = []
    for rule in RULES:
        v = rule(content, path)
        if v:
            violations.append(v)

    return {
        "failures": sum(1 for v in violations if v["severity"] == "FAIL"),
        "warnings": sum(1 for v in violations if v["severity"] == "WARN"),
        "violations": violations,
    }
