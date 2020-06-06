// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by "mdtogo"; DO NOT EDIT.
package commands

var AnnotateShort = `[Alpha] Set an annotation on Resources.`
var AnnotateLong = `
[Alpha]  Set an annotation on Resources.

  DIR:
    Path to local directory.
`
var AnnotateExamples = `
    kustomize cfg annotate my-dir/ --kv foo=bar

    kustomize cfg annotate my-dir/ --kv foo=bar --kv a=b

    kustomize cfg annotate my-dir/ --kv foo=bar --kind Deployment --name foo`

var CatShort = `[Alpha] Print Resource Config from a local directory.`
var CatLong = `
[Alpha]  Print Resource Config from a local directory.

  DIR:
    Path to local directory.
`
var CatExamples = `
    # print Resource config from a directory
    kustomize cfg cat my-dir/

    # wrap Resource config from a directory in an ResourceList
    kustomize cfg cat my-dir/ --wrap-kind ResourceList --wrap-version config.kubernetes.io/v1alpha1 --function-config fn.yaml

    # unwrap Resource config from a directory in an ResourceList
    ... | kustomize cfg cat`

var CompletionShort = `Install shell completion.`
var CompletionLong = `
Install shell completion for kustomize commands and flags -- supports bash, fish and zsh.

    kustomize install-completion

Registers the completion command with known shells (e.g. .bashrc, .bash_profile, etc):

    complete -C /Users/USER/go/bin/kustomize kustomize

Because the completion command is embedded in kustomize directly, there is no need to update
it separately from the kustomize binary.

To uninstall shell completion run:

    COMP_UNINSTALL=1 kustomize install-completion`

var CountShort = `[Alpha] Count Resources Config from a local directory.`
var CountLong = `
[Alpha] Count Resources Config from a local directory.

  DIR:
    Path to local directory.
`
var CountExamples = `
    # print Resource counts from a directory
    kustomize cfg count my-dir/`

var CreateSetterShort = `[Alpha] Create a custom setter for a Resource field`
var CreateSetterLong = `
Create a custom setter for a Resource field by inlining OpenAPI as comments.

  DIR

    A directory containing Resource configuration.

  NAME

    The name of the setter to create.

  VALUE

    The current value of the field, or a substring within the field.
`
var CreateSetterExamples = `
    # create a setter for port fields matching "8080"
    kustomize cfg create-setter DIR/ port 8080 --type "integer" --field port \
         --description "default port used by the app"

    # create a setter for a substring of a field rather than the full field -- e.g. only the
    # image tag, not the full image
    kustomize cfg create-setter DIR/ image-tag v1.0.1 --type "string" \
        --field image --description "current stable release"`

var FmtShort = `[Alpha] Format yaml configuration files.`
var FmtLong = `
[Alpha] Format yaml configuration files.

Fmt will format input by ordering fields and unordered list items in Kubernetes
objects.  Inputs may be directories, files or stdin, and their contents must
include both apiVersion and kind fields.

- Stdin inputs are formatted and written to stdout
- File inputs (args) are formatted and written back to the file
- Directory inputs (args) are walked, each encountered .yaml and .yml file
  acts as an input

For inputs which contain multiple yaml documents separated by \n---\n,
each document will be formatted and written back to the file in the original
order.

Field ordering roughly follows the ordering defined in the source Kubernetes
resource definitions (i.e. go structures), falling back on lexicographical
sorting for unrecognized fields.

Unordered list item ordering is defined for specific Resource types and
field paths.

- .spec.template.spec.containers (by element name)
- .webhooks.rules.operations (by element value)
`
var FmtExamples = `
	# format file1.yaml and file2.yml
	kustomize cfg fmt file1.yaml file2.yml

	# format all *.yaml and *.yml recursively traversing directories
	kustomize cfg fmt my-dir/

	# format kubectl output
	kubectl get -o yaml deployments | kustomize cfg fmt

	# format kustomize output
	kustomize build | kustomize cfg fmt`

var GrepShort = `[Alpha] Search for matching Resources in a directory or from stdin`
var GrepLong = `
[Alpha] Search for matching Resources in a directory or from stdin.

  QUERY:
    Query to match expressed as 'path.to.field=value'.
    Maps and fields are matched as '.field-name' or '.map-key'
    List elements are matched as '[list-elem-field=field-value]'
    The value to match is expressed as '=value'
    '.' as part of a key or value can be escaped as '\.'

  DIR:
    Path to local directory.
`
var GrepExamples = `
    # find Deployment Resources
    kustomize cfg grep "kind=Deployment" my-dir/

    # find Resources named nginx
    kustomize cfg grep "metadata.name=nginx" my-dir/

    # use tree to display matching Resources
    kustomize cfg grep "metadata.name=nginx" my-dir/ | kustomize cfg tree

    # look for Resources matching a specific container image
    kustomize cfg grep "spec.template.spec.containers[name=nginx].image=nginx:1\.7\.9" my-dir/ | kustomize cfg tree`

var InitShort = `[Alpha] Initialize a directory with a Krmfile.`
var InitLong = `
[Alpha]  Initialize a directory with a Krmfile.

  DIR:
    Path to local directory.
`
var InitExamples = `
    # create a Krmfile in the local directory
    kustomize cfg init

    # create a Krmfile in my-dir/
    kustomize cfg init my-dir/`

var ListSettersShort = `[Alpha] List setters for Resources.`
var ListSettersLong = `
List setters for Resources.

  DIR

    A directory containing Resource configuration.

  NAME

    Optional.  The name of the setter to display.
`
var ListSettersExamples = `
  Show setters:

    $ kustomize cfg list-setters DIR/
        NAME      DESCRIPTION   VALUE     TYPE     COUNT   SETBY  
    name-prefix   ''            PREFIX    string   2`

var MergeShort = `[Alpha] Merge Resource configuration files`
var MergeLong = `
[Alpha] Merge Resource configuration files

Merge reads Kubernetes Resource yaml configuration files from stdin or sources packages and write
the result to stdout or a destination package.

Resources are merged using the Resource [apiVersion, kind, name, namespace] as the key.  If any of
these are missing, merge will default the missing values to empty.

Resources specified later are high-precedence (the source) and Resources specified
earlier are lower-precedence (the destination).

For information on merge rules, run:

	kustomize cfg docs merge
`
var MergeExamples = `
    cat resources_and_patches.yaml | kustomize cfg merge > merged_resources.yaml`

var Merge3Short = `[Alpha] Merge diff of Resource configuration files into a destination (3-way)`
var Merge3Long = `
[Alpha] Merge diff of Resource configuration files into a destination (3-way)

Merge3 performs a 3-way merge by applying the diff between 2 sets of Resources to a 3rd set.

Merge3 may be for rebasing changes to a forked set of configuration -- e.g. compute the difference between the original
set of Resources that was forked and an updated set of those Resources, then apply that difference to the fork.

If a field value differs between the ORIGINAL_DIR and UPDATED_DIR, the value from the UPDATED_DIR is taken and applied
to the Resource in the DEST_DIR.

For information on merge rules, run:

	kustomize cfg docs-merge3
`
var Merge3Examples = `
    kustomize cfg merge3 --ancestor a/ --from b/ --to c/`

var RunFnsShort = `[Alpha] Reoncile config functions to Resources.`
var RunFnsLong = `
[Alpha] Reconcile config functions to Resources.

run sequentially invokes all config functions in the directory, providing Resources
in the directory as input to the first function, and writing the output of the last
function back to the directory.

The ordering of functions is determined by the order they are encountered when walking the
directory.  To clearly specify an ordering of functions, multiple functions may be
declared in the same file, separated by '---' (the functions will be invoked in the
order they appear in the file).

#### Arguments:

  DIR:
    Path to local directory.

#### Config Functions:

  Config functions are specified as Kubernetes types containing a metadata.annotations.[config.kubernetes.io/function]
  field specifying an image for the container to run.  This image tells run how to invoke the container.

  Example config function:

	# in file example/fn.yaml
	apiVersion: fn.example.com/v1beta1
	kind: ExampleFunctionKind
	metadata:
	  annotations:
	    config.kubernetes.io/function: |
	      container:
	        # function is invoked as a container running this image
	        image: gcr.io/example/examplefunction:v1.0.1
	    config.kubernetes.io/local-config: "true" # tools should ignore this
	spec:
	  configField: configValue

  In the preceding example, 'kustomize fn run example/' would identify the function by
  the metadata.annotations.[config.kubernetes.io/function] field.  It would then write all Resources in the directory to
  a container stdin (running the gcr.io/example/examplefunction:v1.0.1 image).  It
  would then write the container stdout back to example/, replacing the directory
  file contents.

  See ` + "`" + `kustomize help cfg docs-fn` + "`" + ` for more details on writing functions.
`
var RunFnsExamples = `
kustomize fn run example/`

var SetShort = `[Alpha] Set values on Resources fields values.`
var SetLong = `
Set values on Resources fields.  May set either the complete or partial field value.

` + "`" + `set` + "`" + ` identifies setters using field metadata published as OpenAPI extensions.
` + "`" + `set` + "`" + ` parses both the Kubernetes OpenAPI, as well OpenAPI published inline in
the configuration as comments.

` + "`" + `set` + "`" + ` maybe be used to:

- edit configuration programmatically from the cli
- create reusable bundles of configuration with custom setters

  DIR

    A directory containing Resource configuration.

  NAME

    Optional.  The name of the setter to perform or display.

  VALUE

    Optional.  The value to set on the field.


To print the possible setters for the Resources in a directory, run
` + "`" + `list-setters` + "`" + ` on a directory -- e.g. ` + "`" + `kustomize cfg list-setters DIR/` + "`" + `.

#### Tips

- A description of the value may be specified with ` + "`" + `--description` + "`" + `.
- The last setter for the field's value may be defined with ` + "`" + `--set-by` + "`" + `.
- Create custom setters on Resources, Kustomization.yaml's, patches, etc

The description and setBy fields are left unmodified unless specified with flags.

To create a custom setter for a field see: ` + "`" + `kustomize help cfg create-setter` + "`" + `
`
var SetExamples = `
  Resource YAML: Name Prefix Setter

    # DIR/resources.yaml
    ...
    metadata:
        name: PREFIX-app1 # {"type":"string","x-kustomize":{"setter":[{"name":"name-prefix","value":"PREFIX"}]}}
    ...
    ---
    ...
    metadata:
        name: PREFIX-app2 # {"type":"string","x-kustomize":{"setter":[{"name":"name-prefix","value":"PREFIX"}]}}
    ...

  List setters: Show the possible setters

    $ config list-setters DIR/
        NAME      DESCRIPTION   VALUE     TYPE     COUNT   SETBY  
    name-prefix   ''            PREFIX    string   2

  Perform set: set a new value, owner and description

    $ kustomize cfg set DIR/ name-prefix "test" --description "test environment" --set-by "dev"
    set 2 values

  List setters: Show the new values

    $ config list-setters DIR/
        NAME      DESCRIPTION         VALUE     TYPE     COUNT     SETBY 
    name-prefix   'test environment'   test     string   2          dev

  New Resource YAML:

    # DIR/resources.yaml
    ...
    metadata:
        name: test-app1 # {"description":"test environment","type":"string","x-kustomize":{"setBy":"dev","setter":[{"name":"name-prefix","value":"test"}]}}
    ...
    ---
    ...
    metadata:
        name: test-app2 # {"description":"test environment","type":"string","x-kustomize":{"setBy":"dev","setter":[{"name":"name-prefix","value":"test"}]}}
    ...`

var SinkShort = `[Alpha] Implement a Sink by writing input to a local directory.`
var SinkLong = `
[Alpha] Implement a Sink by writing input to a local directory.

    kustomize fn sink [DIR]

  DIR:
    Path to local directory.  If unspecified, sink will write to stdout as if it were a single file.

` + "`" + `sink` + "`" + ` writes its input to a directory
`
var SinkExamples = `
    kustomize fn source DIR/ | your-function | kustomize fn sink DIR/`

var SourceShort = `[Alpha] Implement a Source by reading a local directory.`
var SourceLong = `
[Alpha] Implement a Source by reading a local directory.

    kustomize fn source DIR...

  DIR:
    One or more paths to local directories.  Contents from directories will be concatenated.
    If no directories are provided, source will read from stdin as if it were a single file.

` + "`" + `source` + "`" + ` emits configuration to act as input to a function
`
var SourceExamples = `
    # emity configuration directory as input source to a function
    kustomize fn source DIR/

    kustomize fn source DIR/ | your-function | kustomize fn sink DIR/`

var TreeShort = `[Alpha] Display Resource structure from a directory or stdin.`
var TreeLong = `
[Alpha] Display Resource structure from a directory or stdin.

kustomize cfg tree may be used to print Resources in a directory or cluster, preserving structure

Args:

  DIR:
    Path to local directory directory.

Resource fields may be printed as part of the Resources by specifying the fields as flags.

kustomize cfg tree has build-in support for printing common fields, such as replicas, container images,
container names, etc.

kustomize cfg tree supports printing arbitrary fields using the '--field' flag.

By default, kustomize cfg tree uses Resource graph structure if any relationships between resources (ownerReferences)
are detected, as is typically the case when printing from a cluster. Otherwise, directory graph structure is used. The
graph structure can also be selected explicitly using the '--graph-structure' flag.
`
var TreeExamples = `
    # print Resources using directory structure
    kustomize cfg tree my-dir/

    # print replicas, container name, and container image and fields for Resources
    kustomize cfg tree my-dir --replicas --image --name

    # print all common Resource fields
    kustomize cfg tree my-dir/ --all

    # print the "foo"" annotation
    kustomize cfg tree my-dir/ --field "metadata.annotations.foo"

    # print the "foo"" annotation
    kubectl get all -o yaml | kustomize cfg tree \
      --field="status.conditions[type=Completed].status"

    # print live Resources from a cluster using owners for graph structure
    kubectl get all -o yaml | kustomize cfg tree --replicas --name --image

    # print live Resources with status condition fields
    kubectl get all -o yaml | kustomize cfg tree \
      --name --image --replicas \
      --field="status.conditions[type=Completed].status" \
      --field="status.conditions[type=Complete].status" \
      --field="status.conditions[type=Ready].status" \
      --field="status.conditions[type=ContainersReady].status"`
