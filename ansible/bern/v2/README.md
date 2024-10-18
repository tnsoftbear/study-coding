```sh
# apply the "init-only" role
ansible-playbook -k init-only.yml
# test the result of the "init-only" role application
ansible-playbook -i init/tests/inventory init/tests/test.yml
```