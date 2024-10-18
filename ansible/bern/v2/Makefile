init-only:		# pw: r
	ansible-playbook -k init-only.yml
mariadb-only:	# pw: a
	ansible-playbook -K mariadb-only.yml
full:			# pw: r
	ansible-playbook -k full.yml
fresh:
	sudo virsh snapshot-revert --domain bern-vm1 --snapshotname 1728015469
test-init:
	ansible-playbook -i init/tests/inventory init/tests/test.yml
test-mariadb:	# pw: a
	ansible-playbook -K -i mariadb/tests/inventory mariadb/tests/test.yml
