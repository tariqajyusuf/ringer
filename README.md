# Ringer: A universal package manager and workstation setup tool

> *ringer 1 | ˈriNGər |*\
> noun\
>   1 **informal** (also dead ringer) a person or thing that looks exactly like
>     another: he is a dead ringer for his late papa\
>   2 a person or device that rings something.\
>   3 a universal package manager for setting up new workstations

Ringer is a universal install command that bridges many of the common platforms.
When setting up new systems, there are usually several pieces of software to
install to get your environment just right. Ringer helps bridge that setup
process across multiple platforms.

## Concepts

### Circle files
Circle files are configuration files that provide context of what you want your
system to be at the end. It is a YAML file that includes the packages,
configurations, and settings that you want on your new system.

### Guise configurations
Guise configurations are definitions of "packages" that provide translation
between the different common platforms. This is how we translate from `ringer
add vscode` to `brew install visual-studio-code` on Mac and `winget install
Microsoft.VisualStudioCode` on Windows.

## Why not puppet/ansible/chef?
A lot of times, particularly within consulting contexts, I have multiple
machines that I'm working with that I cannot remotely manage (or don't want to)
but I want to have the same system setup in different environments. Often these
can be on a range of platforms. I had a simple setup script I was using but this
resulted in having to manage shell scripts that varied on different platforms.

Other remote management tools also assume some uniformity in envrionment and
require you to use different modules depending on what plaform you're
configuring. This was not that far off from what I was doing before and does not
address the non-remote management case.
