

# Terraform Provider for Okta

The Terraform Okta provider is a plugin for Terraform that allows for the full lifecycle management of Okta resources.
This provider is maintained internally by the Okta development team.

## Project structure
 * main.tf will have the okta provider configured
 * csv folder will have all the bulk resources defined in the form of CSV 
 * variables.td will the file where the CSV is decoded into MAP and assigned to terrraform local variables
 * Other Okta resources are defined its its own tf file , like groups.tf, users.tf....
 * scripts folder is the place where all the go lag automation scripts are places
 * logs is the place where all the go lang automation script will log the terraform response


# Running the project 

## Step1
    setting the environment variable is the first step, the env variables need to be overwritten for each tenant you would like to change , an ideal way is to have create multiple activate.sh files for each tenant and set the env before the below steps are achieved 
    
```sh
$ export OKTA_ORG_NAME=dev-60961954(example)
$ export OKTA_BASE_URL=okta.com
$ export OKTA_API_TOKEN=<token>


```

## Step2
Create workspaces by running the following command.
NOTE: you can add more workspaces as you need in the create_workspaces.sh file 
    
```sh
$ ./create_workspaces

```

## Step3
Create workspaces by running the following command.
NOTE: you can add more workspaces as you need in the create_workspaces.sh file 
```sh
$ ./create_workspaces

```

## Step4
All the terraform apply command is automated using a go lang script places under scripts folder 
to apply changes run the following command 
NOTE: the last part of the argument in the command is the workspace

```sh
$ cd scripts
$ go run main.go <workspace>

```