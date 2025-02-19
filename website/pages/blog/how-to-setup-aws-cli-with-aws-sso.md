---
title: How to setup AWS CLI with AWS SSO
tag: security
date: 2022/01/04
description: A step-by-step guide to setup AWS CLI with AWS SSO
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


AWS SSO makes it easy to centrally manage SSO Access to multiple AWS accounts, moves the authentication to the IdP (Identity Provider) and removes the need for managing static, long-lived credentials.

AWS CLI added support for SSO [late 2019](https://aws.amazon.com/blogs/developer/aws-cli-v2-now-supports-aws-single-sign-on/) so you can use it seamlessly in your developer workflow from the CLI without going to the developers portal every time and paste short-lived credentials to the console.


## Prerequisite

### Setup AWS SSO with an IDP

The first step is to have AWS SSO setup and configured. This should be done by someone with the right admin access permissions to both the IdP and AWS. Check out how to set up [AWS SSO with G Suite as IDP](https://www.cloudquery.io/blog/aws-sso-tutorial-with-google-workspace-as-an-idp).

### Install AWS CLI (v2)

On your local machine, if you don’t already have it, install [AWS CLI v2](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html).

## Configure an SSO Profile

Similar to the `aws configure` command that creates a new profile in `~/.aws/config` with long-lived access keys `aws configure sso` command creates a new SSO profile.

`aws configure sso` will prompt you for:

```bash
ep@macbook-pro-73 aws % aws configure sso

# This is the URL that you defined when you setup the AWS
SSO start URL [None]: [https://your-url.awsapps.com/start](https://your-url.awsapps.com/start)

# This is the region that you enabled AWS SSO in
SSO Region [None]: us-east-1

# This step will take you to the browser and you will have to click login and allow
```

![](/images/blog/how-to-setup-aws-cli-with-aws-sso/aws-authorize-request.png)

```bash
# This will suggest to choose an account from which are available to you
There are 6 AWS accounts available to you.

Using the account ID xxxxxxxxxxxx

# This will suggest a role available to you for this account
The only role available to you is: AdministratorAccess

Using the role name "AdministratorAccess"

# Optional: you can choose a default region
CLI default client Region [None]:

# Optional: you can choose a default output form. You can skip this to use the default
CLI default output format [None]:

# Here pick a name that you will be able to use later as an alias for this account for –profile argument
CLI profile name [AdministratorAccess-345990386405]: cq-dev-admin
```

That’s it you configured a new profile (in that case named `cq-dev-admin`) and to test it run the following command:

```bash
aws s3 ls --profile cq-dev-admin
## wil output available s3 buckets
```

## Configure Multiple SSO Profiles

It is common to have multiple accounts available to you via SSO and the neat thing is that you only need to login **once** to any of those **accounts** and you can use any of them in the CLI. The only thing that you will need to do is to add the additional profiles either manually (which will probably be faster) or through the interactive CLI. In either your `~/.aws/config` should look something like the following:

```bash
[profile profile-name-1]
sso_start_url = https://xxxxxx.awsapps.com/start
sso_region = us-east-1
sso_account_id = yyyyyyy
sso_role_name = AdministratorAccess

[profile profile-name-1]
sso_start_url = https://xxxxxx.awsapps.com/start/
sso_region = us-east-1
sso_account_id = yyyyyyy
sso_role_name = AdministratorAccess
region = eu-central-1
```

Once you logged in with any of those profiles (as long as the `sso_start_url` and `sso_region` are the same) with the following command:

```bash
aws sso login --profile profile-name-1
```

you can run also without logging in specifically to other profiles!

```bash
aws s3 ls --profile profile-name-2
```

## Logging out

You can also logout and clear the temporary credentials with `aws sso logout` but this will probably not be necessary most of the time as they expire every hour or so (or a maximum of 12 hours depending on what you admin [defined](https://docs.aws.amazon.com/singlesignon/latest/userguide/howtosessionduration.html) as session duration) so most probably you will have to re-run `aws sso login --profile profile-name-1` once a day or so.

## Summary

If you are using AWS SSO (you probably should :) ) you can definitely enjoy the smooth integration and developer experience with the AWS CLI.

If you are a CloudQuery user you can also enjoy using your temporary SSO credentials in conjunction with CloudQuery seamlessly by specifying AWS_PROFILE=your-sso-profile-name.
