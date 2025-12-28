# Tagitha

A Terraform tagging tool inspired by [Yor](https://github.com/bridgecrewio/yor) and [Checkov](https://github.com/bridgecrewio/checkov)


## Notes

This is my first proper GoLang project. I have tried my best to follow best practice however I am still learning the language.

## History

Within my job I use [Yor](https://github.com/bridgecrewio/checkov) extensively however limitations with how Yor works meant that it started to become unsuitable for what we wanted it to do, so I started to maintain [my own version](https://github.com/flamableassassin/yor).
However I wasn't completely satisfied with it and so I started to look for alternatives such as [TerraTag](https://github.com/env0/terratag) however it doesn't offer features such as [Git tagging](https://github.com/env0/terratag/issues/212).
So I slept on the idea for a fair few months and here we are...Tagitha _(Taggy was [taken already](https://github.com/open-taggy/taggy))_


## Features

- [x] More complex filtering system.
  - [ ] Directory based
  - [ ] Existing tags
  - [ ] Resource type
  - [ ] Resource name
  - [ ] File name
  - [ ] *Attribute based tagging via templating system
- [ ] Git info tagging
- [ ] Ignore resources with a comment
- [ ] *Modify the tag names globally and per group.
    - Default < Global < Group  

_Items with a `*` are ideas_