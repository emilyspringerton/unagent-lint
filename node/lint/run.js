const fs = require('fs');
const rules = require('./rules');

function run(path) {
  const content = fs.readFileSync(path, 'utf8');
  const violations = [];

  rules.forEach(rule => {
    const v = rule.check(content, path);
    if (v) violations.push(v);
  });

  return {
    failures: violations.filter(v => v.severity === "FAIL").length,
    warnings: violations.filter(v => v.severity === "WARN").length,
    violations
  };
}

module.exports = { run };
