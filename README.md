# pruneOldOpenstackImages

pruneOldOpenstackImages is tool to delete duplicate OpenStack images.

## Use

OpenStack (as many other computing platforms) allows you to create multiple VM images with the same name. This usually happens when you create a new version of the image after operations like update the OS packages. Image creation tools (such as Packer) are not resposible to clear up these duplicate images. So, a process needs to be implemented to clear them up. You can use pruneOldOpenstackImages to achieve this goal.

It is a good idea to have it as a part of image creation process in your CI flow and run at the end to prune unwanted, old duplicate images.

## Terms of USE

pruneOldOpenstackImages is being provided "AS IS".  You acknowledge that the "pruneOldOpenstackImages" is not error-free. You use it at your own risk and discretion. That means the "pruneOldOpenstackImages" doesn’t come with any warranty. None express, none implied. The "pruneOldOpenstackImages" will be continually developed, and you acknowledge that changes to functionality and layout may carried out without advance notice. 

## Installation

## Build from Source Code

First please modify Makefile to be sure if your Operating System and Architecture is correctly set. In order to compile against Linux TARGET_OS should be set to linux

Then, issue the command "make", it should compile the source code and produce a binary named pruneOldOpenstackImages

You can use the "static" target to produce a statically build binary, if you'd like.

## Usage

```
pruneOldOpenstackImages has several command line options, the can be seen by issuing a pruneOldOpenstackImages -h command.

usage: pruneOldOSImages --imageName=IMAGENAME --authFile=AUTHFILE [<flags>]

Flags:
  -h, --help                 Show context-sensitive help (also try --help-long and --help-man).
      --check                Enable check mode, don't actually delete anything
      --imageName=IMAGENAME  Name of the image to save
      --numImages=2          Number of images with same name to keep
      --region=REGION        Region
      --authFile=AUTHFILE    Absolute path of a JSON file that contains the authentication information
```

## Auth File

pruneOldOpenstackImages needs a json file to read authentication data from. You can find an example of this json file in this repository.

## Authors

      Ozgur Demir <ozgurcd@gmail.com>
