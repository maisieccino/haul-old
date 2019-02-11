# Haul Shared Doc Format

## How Haul does storage management

YAML files. lots of YAML files

* Thing - a block of something. Rich text (markdown), image, nested thing (lists of text, todo lists), and so on...
 - Can be assigned a 'due date'/'reminder date'. For reminders etc
* Page - a collection of things. Top level thing defines what this page is. It might be a to-do list, or a Kanban board, or...
 - A page is actually a thing.
* Bag - a place to store pages. Kinda like folders, except a bit more general. Basically like an S3 bucket. Good for namespacing stuff
* Label - a way of grouping stuff under topics. These can be given emoji, colour, etc, even a default page.

## Storage directory structure
Directory structure of `/things` to match that of the bags. just to make RBAC stuff easier.

```
things/
  <bag/path>/
    _bag.yaml
    <date-created>-<random-hash>.yaml
labels/
  <label_name>.yaml
extensions/
  types/
    <type>.yaml
haul.yaml
things.yaml
```

## `_bag.yaml`

Defines the access rules for the bag.

## `<label>.yaml`
Defines a label named `<label>`.

```yaml
apiVersion: haul.io/v1alpha1
kind: label
metadata:
  dateCreated: 2018-12-14T20:45:00Z
  dateModified: 2018-12-14T20:45:00Z
  humanName: 'My Label'
  emoji: '🆒'
  color: '#ff0000'
spec:
  page: {}
```

## `haul.yaml`

Defines a Haul.

```yaml
apiVersion: haul.io/v1alpha1
kind: haul
metadata:
  dateCreated: 2018-12-14T20:45:00Z
  dateModified: 2018-12-14T20:45:00Z
  owner: 'Matt'
  thingCount: 22
spec:
  pinnedPages:
    - thingRef: 235acb3b1
    - thingRef: 235acb3b1
```

## `<date-created>-<random-hash>.yaml`

A file that defines a _thing_. e.g.

```yaml
apiVersion: thing.haul.io/v1alpha1
kind: richText
metadata:
  dateCreated: 2018-12-14T20:45:00Z
  dateModified: 2018-12-14T20:45:00Z
  owner: 'Matt'
  bag: bag/name
  labels:
  - label-name
  - other-label
spec:
  content:
    base64: |
      This is some random rich text that I've added here. Supports markdown etc.
      Also can do neat things like reference other things:

      @{thing:2018-10-10T120000Z-abcdef123}

      This might import an image, for example.
```

```yaml
apiVersion: thing.haul.io/v1alpha1
kind: image
metadata:
  dateCreated: 2018-12-14T20:45:00Z
  dateModified: 2018-12-14T20:45:00Z
  owner: 'Matt'
  bag: bag/name
  labels:
  - label-name
  - other-label
spec:
  path: /storage/images/<image>.jpg
  altText: This is a photo of a pretty flower.
```
