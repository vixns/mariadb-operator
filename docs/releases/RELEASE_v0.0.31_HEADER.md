This is a patch release to fix some regressions related to the Galera. If you are using Galera and `ReadWriteOnce` storage, it is highly recommended to upgrade.

When using `ReadWriteOnce` storage, if the Galera recovery `Jobs` (introduced in [v0.0.30](https://github.com/mariadb-operator/mariadb-operator/releases/tag/v0.0.30)) are not scheduled in the same node as the `MariaDB` `Pods`, they will be unable to attach the PVC and the recovery will not progress. See the following issues for further detail: https://github.com/mariadb-operator/mariadb-operator/issues/840 https://github.com/mariadb-operator/mariadb-operator/issues/848

By default, we have configured `podAffinity` to the recovery `Jobs` to make sure the attachment succeeds. See the PR: https://github.com/mariadb-operator/mariadb-operator/pull/843

In addition, we have also fixed another regression that prevents the operator from  polling after the related `MariaDB` resource has been deleted. See the PR: https://github.com/mariadb-operator/mariadb-operator/pull/832

Finally, to better support large fleets of databases, we are introducing support for `MaxConcurrentReconciles`, defaulting to `10` for `MariaDB` and `MaxScale` resources: https://github.com/mariadb-operator/mariadb-operator/pull/833

---

We value your feedback! If you encounter any issues or have suggestions, please [open an issue on GitHub](https://github.com/mariadb-operator/mariadb-operator/issues/new/choose). Your input is crucial to improve `{{ .ProjectName }}`🦭.

Join us on Slack: **[MariaDB Community Slack](https://r.mariadb.com/join-community-slack)**.