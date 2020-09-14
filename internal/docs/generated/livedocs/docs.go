// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package livedocs

var LiveShort = `Reconcile configuration files with the live state`
var LiveLong = `
| Reads From              | Writes To                |
|-------------------------|--------------------------|
| local files             | cluster                  |
| cluster                 | stdout                   |

Live contains the next-generation versions of apply related commands for
deploying local configuration packages to a cluster.
`

var ApplyShort = `Apply a package to the cluster (create, update, delete)`
var ApplyLong = `
  kpt live apply DIR [flags]

Args:

  DIR:
    Path to a package directory.  The directory must contain exactly
    one ConfigMap with the grouping object annotation.

Flags:

  --poll-period:
    The frequency with which the cluster will be polled to determine
    the status of the applied resources. The default value is every 2 seconds.
  
  --reconcile-timeout:
    The threshold for how long to wait for all resources to reconcile before
    giving up. If this flag is not set, kpt live apply will not wait for
    resources to reconcile.
  
  --prune-timeout:
    The threshold for how long to wait for all pruned resources to be
    deleted before giving up. If this flag is not set, kpt live apply will not
    wait. In most cases, it would also make sense to set the
    --prune-propagation-policy to Foreground when this flag is set.
  
  --prune-propagation-policy:
    The propagation policy kpt live apply should use when pruning resources. The
    default value here is Background. The other options are Foreground and Orphan.
  
  --output:
    This determines the output format of the command. The default value is
    events, which will print the events as they happen. The other option is
    table, which will show the output in a table format.
`
var ApplyExamples = `
  # apply resources and prune
  kpt live apply my-dir/

  # apply resources and wait for all the resources to be reconciled before pruning
  kpt live apply --reconcile-timeout=15m my-dir/

  # apply resources and specify how often to poll the cluster for resource status
  kpt live apply --reconcile-timeout=15m --poll-period=5s my-dir/
`

var DestroyShort = `Remove all previously applied resources in a package from the cluster`
var DestroyLong = `
  kpt live destroy DIR

Args:

  DIR:
    Path to a package directory.  The directory must contain exactly
    one ConfigMap with the grouping object annotation.
`
var DestroyExamples = `
  # remove all resources in a package from the cluster
  kpt live destroy my-dir/
`

var DiffShort = `Diff the local package config against the live cluster resources`
var DiffLong = `
  kpt live diff DIR
  
  Output is always YAML.
  
  KUBECTL_EXTERNAL_DIFF environment variable can be used to select your own diff command. By default, the "diff" command
  available in your path will be run with "-u" (unicode) and "-N" (treat new files as empty) options.

Args:

  DIR:
    Path to a package directory.  The directory must contain exactly one ConfigMap with the inventory annotation.

Exit Status:

  The following exit values shall be returned:
  
  0 No differences were found. 1 Differences were found. >1 kpt live or diff failed with an error.
  
  Note: KUBECTL_EXTERNAL_DIFF, if used, is expected to follow that convention.
`
var DiffExamples = `
  # diff the config in "my-dir" against the live cluster resources
  kpt live diff my-dir/
  
  # specify the local diff program to use
  export KUBECTL_EXTERNAL_DIFF=meld; kpt live diff my-dir/
`

var FetchK8sSchemaShort = `Fetch the OpenAPI schema from the cluster`
var FetchK8sSchemaLong = `
  kpt live fetch-k8s-schema [flags]

Flags:

  --pretty-print
    Format the output before printing
`
var FetchK8sSchemaExamples = `
  # print the schema for the cluster given by the current context
  kpt live fetch-k8s-schema
  
  # print the schema after formatting using a named context
  kpt live fetch-k8s-schema --context=myContext --pretty-print
`

var InitShort = `Initialize a package with a object to track previously applied resources`
var InitLong = `
  kpt live init DIRECTORY [flags]

Args:

  DIR:
    Path to a package directory.  The directory must contain exactly
    one ConfigMap with the grouping object annotation.

Flags:

  --inventory-id:
    Identifier for group of applied resources. Must be composed of valid label characters.
  --inventory-namespace:
    namespace for the inventory object. If no namespace is provided, namespace "default" will be used.
`
var InitExamples = `
  # initialize a package
  kpt live init my-dir/

  # initialize a package with a specific name for the group of resources
  kpt live init --inventory-namespace=test my-dir/
`

var PreviewShort = `Preview prints the changes apply would make to the cluster`
var PreviewLong = `
  kpt live preview DIRECTORY [flags]

Args:

  DIRECTORY:
    One directory that contain k8s manifests. The directory
    must contain exactly one ConfigMap with the grouping object annotation.

Flags:

  --destroy:
    If true, dry-run deletion of all resources.
`
var PreviewExamples = `
  # preview apply for a package
  kpt live preview my-dir/

  # preview destroy for a package
  kpt live preview --destroy my-dir/
`

var StatusShort = `Status shows the status for the resources in the cluster`
var StatusLong = `
  kpt live status (DIR | STDIN) [flags]

Args:

  DIR | STDIN:
    Path to a directory if an argument is provided or reading from stdin if left
    blank. In both situations one of the manifests must contain exactly one
    ConfigMap with the inventory template annotation.

Flags:

  --poll-period (duration):
    The frequency with which the cluster will be polled to determine the status
    of the applied resources. The default value is every 2 seconds.
  
  --poll-until (string):
    When to stop polling for status and exist. Must be one of the following:
      known:   Exit when the status for all resources have been found.
      current: Exit when the status for all resources have reached the Current status.
      deleted: Exit when the status for all resources have reached the NotFound
               status, i.e. all the resources have been deleted from the live state.
      forever: Keep polling for status until interrupted.
    The default value is ‘known’.
  
  --output (string):
    Determines the output format for the status information. Must be one of the following:
      events: The output will be a list of the status events as they become available.
      table:  The output will be presented as a table that will be updated inline
              as the status of resources become available.
    The default value is ‘events’.
  
  --timeout (duration):
    Determines how long the command should run before exiting. This deadline will
    be enforced regardless of the value of the --poll-until flag. The default is
    to wait forever.
`
var StatusExamples = `
  # Monitor status for a set of resources based on manifests. Output in table format:
  kpt live status my-app/ --poll-until=forever --output=table

  # Check status for a set of resources read from stdin with output in events format
  kpt cfg cat my-app | kpt live status
`
