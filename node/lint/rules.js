module.exports = [
  {
    id: "ID-001",
    severity: "FAIL",
    check(content, file) {
      const lowered = content.toLowerCase();
      if (lowered.includes("acts on your behalf")) {
        return {
          rule: "ID-001",
          severity: "FAIL",
          message: "UNAGENT must not be described as an autonomous agent.",
          file
        };
      }
    }
  }
];
