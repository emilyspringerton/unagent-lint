def not_an_agent(content, file):
    if "acts on your behalf" in content.lower():
        return {
            "rule": "ID-001",
            "severity": "FAIL",
            "message": "UNAGENT must not be described as an autonomous agent.",
            "file": file,
        }


RULES = [not_an_agent]
