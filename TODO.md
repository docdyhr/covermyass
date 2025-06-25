# TODO - Covermyass Improvements

This document outlines recommended improvements for the covermyass project, organized by priority level.

## 🔴 HIGH PRIORITY

### Code Quality & Bug Fixes

- [ ] **Fix failing shred tests** (`lib/shred/shred_test.go`)
  - Test logic uses `assert.NoError` but expects errors
  - Should use `assert.Error` for error test cases
  - Create missing `lib/shred/testdata/` directory with required test files

- [ ] **Improve concurrent error handling** (`lib/analysis/analyzer.go:38-42`)
  - Replace simple error logging with proper error aggregation
  - Implement error collection from goroutines
  - Add context cancellation on first error

### Security Concerns

- [ ] **Add file permission verification** (`lib/shred/shred.go`)
  - Verify write permissions before attempting shred operations
  - Add user confirmation for system-critical files
  - Implement safeguards against accidental system file deletion

- [ ] **Enhance random data generation** (`lib/shred/shred.go:88`)
  - Ensure crypto/rand is used consistently
  - Add entropy verification for random data
  - Consider using multiple entropy sources

- [ ] **Implement system file protection**
  - Add blacklist of critical system files (e.g., `/boot/*`, `/etc/passwd`)
  - Create warning system for potentially dangerous operations
  - Add `--force` flag for overriding protections

### Missing Critical Functionality

- [ ] **Add dry-run mode** (`cmd/root.go`)
  - Implement `--dry-run` flag for safe operation testing
  - Show what would be deleted without actually deleting
  - Include size estimates and operation time predictions

- [ ] **Implement backup mechanism**
  - Add optional backup creation before shredding
  - Secure deletion of backups after confirmation
  - Backup integrity verification

## 🟡 MEDIUM PRIORITY

### Testing Improvements

- [ ] **Complete finder package tests** (`lib/find/finder_test.go`)
  - Implement comprehensive finder tests with mock filesystem
  - Test glob pattern matching edge cases
  - Add performance tests for large directory structures

- [ ] **Add integration tests**
  - Create end-to-end test suite covering complete workflow
  - Test cross-platform compatibility
  - Add performance benchmarks

- [ ] **Expand test coverage**
  - Add tests for `output/printer.go`
  - Increase error handling test coverage
  - Add property-based testing for file operations

### Architecture Improvements

- [ ] **Refactor check registration system** (`lib/check/checks.go`)
  - Replace global state with dependency injection
  - Create check registry interface
  - Enable runtime check registration/deregistration

- [ ] **Decouple analysis and output** (`lib/analysis/analyzer.go`)
  - Separate output formatting from analysis logic
  - Create pluggable output formatters (JSON, XML, CSV)
  - Improve testability of analysis logic

- [ ] **Design plugin system for checks**
  - Create plugin interface for custom check implementations
  - Add plugin discovery and loading mechanism
  - Document plugin development guidelines

### Performance Optimizations

- [ ] **Optimize memory usage for large files** (`lib/shred/shred.go:86`)
  - Implement chunked reading/writing instead of loading entire files
  - Add memory usage monitoring and limits
  - Progressive memory allocation based on file size

- [ ] **Improve concurrent operations** (`lib/analysis/analyzer.go`)
  - Implement worker pool pattern to limit concurrent goroutines
  - Add configurable concurrency limits
  - Optimize resource usage for large-scale operations

### Cross-Platform Compatibility

- [ ] **Complete Windows support**
  - Finish Windows implementation (currently commented out in build config)
  - Add Windows-specific log file patterns
  - Test Windows file permission handling

- [ ] **Standardize path handling** (`lib/find/finder.go:49-52`)
  - Use `filepath.Join` consistently across all platforms
  - Fix path separator handling inconsistencies
  - Add path validation for different filesystems

## 🟢 LOW PRIORITY

### Documentation & Usability

- [ ] **Expand developer documentation**
  - Add comprehensive API documentation
  - Create contribution guidelines
  - Document architecture decisions and design patterns

- [ ] **Improve CLI help and examples** (`cmd/root.go`)
  - Add more detailed usage examples
  - Include safety warnings in help text
  - Create interactive help system

### Development Workflow

- [ ] **Add pre-commit hooks**
  - Implement linting, testing, and security checks
  - Add commit message validation
  - Integrate with existing development tools

- [ ] **Enhance CI/CD pipeline**
  - Add comprehensive security scanning
  - Implement automated dependency updates
  - Add performance regression testing

### Configuration Management

- [ ] **Add configuration file support**
  - Support for custom check patterns via config files
  - User preference storage (default iterations, filters)
  - Environment-specific configuration profiles

- [ ] **Make check patterns configurable**
  - Extract hard-coded patterns from check files
  - Create external pattern definition format
  - Enable user-defined check additions

## 🛡️ Security Enhancements

- [ ] **File integrity verification**
  - Verify file integrity before and after operations
  - Implement checksums for verification
  - Add corruption detection mechanisms

- [ ] **Audit logging**
  - Log all file operations with timestamps
  - Include user context and operation details
  - Support for external log forwarding

- [ ] **User permission validation**
  - Validate user privileges before operations
  - Implement role-based operation restrictions
  - Add privilege escalation warnings

## 📈 Performance & Monitoring

- [ ] **Progress reporting enhancements**
  - Add granular progress reporting options
  - Implement ETA calculations
  - Support for quiet/verbose modes

- [ ] **Memory usage monitoring**
  - Add memory usage limits and monitoring
  - Implement automatic memory cleanup
  - Add memory usage reporting

- [ ] **File size-based strategies**
  - Implement different strategies for small vs large files
  - Optimize operations based on file characteristics
  - Add file type-specific handling

## 🔧 Error Handling Improvements

- [ ] **Structured error types**
  - Create specific error types for different failure modes
  - Implement error recovery mechanisms
  - Add detailed error context information

- [ ] **Partial failure recovery**
  - Handle partial operation failures gracefully
  - Implement operation rollback mechanisms
  - Add retry logic for transient failures

- [ ] **User-friendly error reporting**
  - Provide clear error messages with suggested solutions
  - Add troubleshooting guides for common issues
  - Implement error classification system

---

## Implementation Notes

- **Security First**: Always prioritize security-related improvements
- **Testing**: Each improvement should include corresponding tests
- **Documentation**: Update documentation alongside code changes
- **Backwards Compatibility**: Consider impact on existing users
- **Performance**: Profile before and after performance-related changes

## Getting Started

1. Review the HIGH PRIORITY items first
2. Fix failing tests to establish a stable testing baseline
3. Address security concerns before adding new features
4. Consider creating GitHub issues for tracking progress on larger items