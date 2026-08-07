# Security Policy

Thank you for helping keep Kavrok and its users secure.

We take security issues seriously and appreciate responsible disclosure of vulnerabilities.

---

# Supported Versions

The following versions of Kavrok currently receive security updates.

| Version | Supported |
|----------|-----------|
| Latest Release | ✅ |
| Previous Minor Release | ✅ |
| Older Releases | ❌ |
| Development Branch (`main`) | Best Effort |

Only the latest stable release and the previous minor release are guaranteed to receive security fixes.

---

# Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, report them privately by emailing:

**kasab.mohammed.imran@gmail.com**

Please include as much information as possible:

- Description of the vulnerability
- Affected version(s)
- Steps to reproduce
- Proof of concept (if available)
- Impact assessment
- Suggested remediation (optional)

Providing detailed information helps us investigate and resolve issues more quickly.

---

# What to Expect

After receiving your report, we aim to:

- Acknowledge receipt within **3 business days**
- Provide an initial assessment within **7 business days**
- Keep you informed about progress
- Release a fix as quickly as practical

Complex vulnerabilities may require additional investigation and coordination.

---

# Responsible Disclosure

We kindly ask that you:

- Give us reasonable time to investigate and fix the issue before public disclosure.
- Avoid accessing, modifying, or deleting data that does not belong to you.
- Avoid actions that could negatively impact users or systems.
- Do not exploit vulnerabilities beyond what is necessary to demonstrate the issue.

We are committed to working collaboratively with security researchers and contributors.

---

# Scope

This policy applies to:

- Kavrok CLI
- Official source code
- Official releases
- Documentation
- Build and release workflows

Third-party dependencies are outside the scope of this policy and should be reported to their respective maintainers.

---

# Security Best Practices

When using Kavrok, we recommend:

- Always use the latest stable release.
- Keep Kubernetes clusters up to date.
- Review generated recommendations before applying them in production.
- Follow the principle of least privilege for Kubernetes credentials.
- Avoid running Kavrok with unnecessary administrative privileges.

---

# Dependency Security

Kavrok regularly updates its dependencies and performs automated dependency checks where possible.

Security-related dependency updates may be released independently of feature releases.

---

# Security Advisories

When appropriate, security fixes will be published through:

- GitHub Security Advisories
- GitHub Releases
- CHANGELOG.md

---

# Acknowledgements

We appreciate the efforts of security researchers and contributors who help improve the security of Kavrok through responsible disclosure.

Thank you for helping make Kavrok more secure.