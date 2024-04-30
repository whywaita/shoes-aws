# shoes-aws: shoes provider for [Amazon Web Services](https://aws.amazon.com)

shoes-aws is [myshoes](https://github.com/whywaita/myshoes) provider

## Setup

Please set environment values.

This project uses LocalStack which mocks AWS services

LocalStack has requirements about Auth - https://docs.localstack.cloud/references/credentials/
- specifically, `test` is a good option for the environment of running test
  > AWS_ACCESS_KEY_ID=test
  >
  > AWS_SECRET_ACCESS_KEY=DOESNOTMATTER


### Required

- `AWS_RESOURCE_TYPE_MAPPING`
    - mapping from [resource_type](https://github.com/whywaita/myshoes/blob/master/docs/how-to-develop-shoes.md#resource-type) to instance type of AWS.
    - e.g.) `{"nano": "c5a.large", "micro": "c5a.xlarge"}`
- Credential values for AWS
    - AWS Shared Configuration
    - See [official documents](https://docs.aws.amazon.com/sdkref/latest/guide/creds-config-files.html)

### Optional

- `AWS_IMAGE_ID`
    - AMI ID for runner
    - default: `ami-02868af3c3df4b3aa`
