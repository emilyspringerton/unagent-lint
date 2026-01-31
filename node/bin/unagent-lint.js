#!/usr/bin/env node
const { run } = require('../lint/run');

const path = process.argv[2];
if (!path) {
  console.error("Path required");
  process.exit(3);
}

const report = run(path);
if (report.failures) process.exit(1);
if (report.warnings) process.exit(2);
