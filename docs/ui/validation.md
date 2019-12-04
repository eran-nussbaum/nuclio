# UI field validation

## Overview

This document describes the different restrictions for various input fields of function configuration in Nuclio UI.

## Dict of RegEx patterns used in this document

| Key  | Value  | Reference/source |
| :--- | :----- | :-------- |
| `configMapKey` | `^(?=[\S\s]{1,253}$)(?!\.$)(?!\.\.[\S\s]*$)[-._a-zA-Z0-9]+$` | [K8s](https://github.com/kubernetes/apimachinery/blob/master/pkg/util/validation/validation.go#L375) |
| `container` | `^(?!.*--)(?!.*__)(?=.*[a-z])[a-z0-9][a-z0-9-_]*[a-z0-9]$\|^[a-z]$` | [Iguazio platform](https://github.com/iguazio/zebo/blob/development/py/services/container_provisioning/__init__.py#L835) |
| `dns1123Label` | `^(?=[\S\s]{1,63}$)[a-z0-9]([-a-z0-9]*[a-z0-9])?$` | [K8s](https://github.com/kubernetes/apimachinery/blob/master/pkg/util/validation/validation.go#L134) |
| `dns1123Subdomain` | `^(?=[\S\s]{1,253}$)[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$` | [K8s](https://github.com/kubernetes/apimachinery/blob/master/pkg/util/validation/validation.go#L155) |
| `dockerReference` | `^(([a-zA-Z0-9]\|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])(\.([a-zA-Z0-9]\|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]))*(\:\d+)?\/)?[a-z0-9]+(([._]\|__\|[-]*)[a-z0-9]+)*(\/[a-z0-9]+(([._]\|__\|[-]*)[a-z0-9]+)*)*(\:[\w][\w.-]{0,127})?(\@[A-Za-z][A-Za-z0-9]*([-_+.][A-Za-z][A-Za-z0-9]*)*\:[0-9a-fA-F]{32,})?$` | [Docker](https://github.com/docker/distribution/blob/master/reference/regexp.go) |
| `envVarName` | `^(?!\.$)(?!\.\.[\S\s]*$)[-._a-zA-Z][-._a-zA-Z0-9]*$` | [K8s](https://github.com/kubernetes/apimachinery/blob/master/pkg/util/validation/validation.go#L359) |
| `prefixedQualifiedName` | `^(?!kubernetes.io\/)(?!k8s.io\/)((?=[\S\s]{1,253}\/)([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*\/))?(?=[\S\s]{1,63}$)([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$` | [K8s](https://github.com/kubernetes/apimachinery/blob/master/pkg/util/validation/validation.go#L42) |
| `qualifiedName` | `^(?=[\S\s]{1,63}$)([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$` | [K8s](https://github.com/kubernetes/apimachinery/blob/master/pkg/util/validation/validation.go#L36) |
| `wildcardDns1123Subdomain` | `^(?=[\S\s]{1,253}$)\*\.[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$` | [K8s](https://github.com/kubernetes/apimachinery/blob/master/pkg/util/validation/validation.go#L201) |

## Function Configuration Input Validation

<table>
<tr align="left">
  <th style="vertical-align:'top'; font-weight:bold;">
    Panel
  </th>
  <th style="vertical-align:'top'; font-weight:bold;">
    Field
  </th>
  <th style="vertical-align:'top'; font-weight:bold;">
    Is Unique?
  </th>
  <th style="vertical-align:'top'; font-weight:bold;">
    Validation Pattern
  </th>
  <th style="vertical-align:'top'; font-weight:bold;">
    Max Length
  </th>
  <th style="vertical-align:'top'; font-weight:bold;">
    Hover Tooltip
  </th>
</tr>
<tr id="annotation-key">
  <td style="vertical-align:'top';">Annotations</td>
  <td>Key</td>
  <td>Yes</td>
  <td><code>prefixedQualifiedName</code></td>
  <td>317 chars</td>
  <td>Annotation keys are composed of an optional prefix and a name, separated by a forward slash (/) &mdash; "&lt;key prefix&gt;/&lt;key name&gt;".
    <br/><br/>
    Name restrictions
    <ul>
      <li>Valid characters &mdash;
        <ul>
          <li>Alphanumeric characters (a&ndash;z, A&ndash;Z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
          <li>Underscores (_)
          </li>
          <li>Periods (.)
          </li>
        </ul>
      </li>
      <li>Must begin and end with alphanumeric characters
      </li>
      <li>Max length &mdash; 63 characters
      </li>
    </ul>
    Prefix restrictions
    <ul>
      <li>Valid characters &mdash;
        <ul>
          <li>Lowercase alphanumeric characters (a&ndash;z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
          <li>Periods (.)
          </li>
        </ul>
      <li>Must begin and end with lowercase alphanumeric characters  (a&ndash;z, 0&ndash;9)
      </li>
      <li>Max length &mdash; 253 characters
      </li>
    </ul>
    Examples
    <ul>
      <li>"MyName"
      </li>
      <li>"sub-domain.example.com/MyName"
      </li>
      <li>"my.name_123"
      </li>
      <li>"123-abc"
      </li>
    </ul>
  </td>
</tr>
<tr id="annotation-value">
  <td>Annotations</td>
  <td>Value</td>
  <td>No</td>
  <td align="center">&ndash;</td>
  <td align="center">&ndash;</td>
  <td align="center">&ndash;</td>
</tr>
<tr id="build-image-name">
  <td>Build</td>
  <td>Image name</td>
  <td>No</td>
  <td><code>dockerReference</code></td>
  <td>255 chars</td>
  <td align="center">&ndash;</td>
</tr>
<tr id="envar-key">
  <td>Environment Variables</td>
  <td>Key</td>
  <td>Yes</td>
  <td><code>envVarName</code></td>
  <td align="center">&ndash;</td>
  <td>Restrictions
    <ul>Valid characters &mdash;
      <li>
        <ul>
          <li>Alphanumeric characters (a&ndash;z, A&ndash;Z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
          <li>Underscores (_)
          </li>
          <li>Periods (.)
          </li>
        </ul>
      </li>
      <li>Must not start with a digit (0&ndash;9) or with two periods (..)
      </li>
      <li>Must not end with a digit (0&ndash;9)
      </li>
      <li>Must not contain only a period (.)
      </li>
    </ul>
    Examples
    <ul>
      <li>"MY_ENV_VAR"</li>
      <li>"MyEnvVar1"</li>
      <li>"My-Env-Var.1"</li>
      <li>"my.env-var"</li>
    </ul>
  </td>
</tr>
<tr id="envar-configmap-key">
  <td>Environment Variables (ConfigMap)</td>
  <td>ConfigMap Key</td>
  <td>Yes</td>
  <td><code>configMapKey</code></td>
  <td></td>
  <td>Restrictions
    <ul>
      <li>Valid characters &mdash;
        <ul>
          <li>Alphanumeric characters (a&ndash;z, A&ndash;Z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
          <li>Underscores (_)
          </li>
        </ul>
      </li>
      <li>Max length &mdash; 253 characters
      </li>
    </ul>
    Examples
    <ul>
      <li>"MY_KEY"</li>
      <li>"my-key"</li>
      <li>"MyKey.1"</li>
    </ul>
  </td>
</tr>
<tr id="envar-value">
  <td>Environment Variables</td>
  <td>Value</td>
  <td>No</td>
  <td align="center">&ndash;</td>
  <td align="center">&ndash;</td>
  <td align="center">&ndash;</td>
</tr>
<tr id="label-key">
  <td style="vertical-align:'top';">Labels</td>
  <td>Key</td>
  <td>Yes</td>
  <td><code>prefixedQualifiedName</code></td>
  <td>317 chars</td>
  <td>Label keys are composed of an optional prefix and a name, separated by a forward slash (/) &mdash; "&lt;key prefix&gt;/&lt;key name&gt;".
    <br/><br/>
    Name restrictions
    <ul>
      <li>Valid characters &mdash;
        <ul>
          <li>Alphanumeric characters (a&ndash;z, A&ndash;Z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
          <li>Underscores (_)
          </li>
          <li>Periods (.)
          </li>
        </ul>
      </li>
      <li>Must begin and end with alphanumeric characters
      </li>
      <li>Max length &mdash; 63 characters
      </li>
    </ul>
    Prefix restrictions
    <ul>
      <li>Valid characters &mdash;
        <ul>
          <li>Lowercase alphanumeric characters (a&ndash;z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
          <li>Periods (.)
          </li>
        </ul>
      <li>Must begin and end with lowercase alphanumeric characters  (a&ndash;z, 0&ndash;9)
      </li>
      <li>Max length &mdash; 253 characters
      </li>
    </ul>
    Examples
    <ul>
      <li>"MyName"
      </li>
      <li>"sub-domain.example.com/MyName"
      </li>
      <li>"my.name_123"
      </li>
      <li>"123-abc"
      </li>
    </ul>
  </td>
</tr>
<tr id="labels-value">
  <td>Labels</td>
  <td>Value</td>
  <td>No</td>
  <td><code>qualifiedName</code></td>
  <td>63 chars</td>
  <td>Restrictions
    <ul>
      <li>Valid characters &mdash;
        <ul>
          <li>Alphanumeric characters (a&ndash;z, A&ndash;Z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
          <li>Underscores (_)
          </li>
          <li>Periods (.)
          </li>
        </ul>
      </li>
      <li>Must begin and end with alphanumeric characters
      </li>
      <li>Max length &mdash; 63 characters
      </li>
    </ul>
    Examples
    <ul>
      <li>"MyValue"</li>
      <li>"my_value.1"</li>
      <li>"12345"</li>
    </ul>
  </td>
</tr>
<tr id="volumes-name">
  <td>Volumes</td>
  <td>Name</td>
  <td>Yes</td>
  <td><code>dns1123Label</code></td>
  <td>63 chars</td>
  <td>Restrictions
    <ul>
      <li>Valid characters &mdash;
        <ul>
          <li>Lowercase alphanumeric characters (a&ndash;z, 0&ndash;9)
          </li>
          <li>Hyphens (-)
          </li>
        </ul>
      <li>Must begin and end with lowercase alphanumeric characters  (a&ndash;z, 0&ndash;9)
      </li>
      <li>Max length &mdash; 63 characters
      </li>
    </ul>
    Examples
    <ul>
      <li>"my_volume"</li>
      <li>"123-abc"</li>
    </ul>
  </td>
</tr>
<tr id="volumes-mount-path">
  <td>Volumes</td>
  <td>Mount Path</td>
  <td>Yes</td>
  <td align="center">&ndash;</td>
  <td align="center">&ndash;</td>
  <td align="center">&ndash;</td>
</tr>
<tr id="volumes-v3io-container-name">
  <td>Volumes (V3IO)</td>
  <td>Container Name</td>
  <td>No</td>
  <td><code>container</code></td>
  <td>128 chars</td>
  <td align="center">&ndash;</td>
</tr>
<tr id="volumes-v3io-sub-path">
  <td>Volumes (V3IO)</td>
  <td>Sub Path</td>
  <td>No</td>
  <td align="center">&ndash;</td>
  <td>255 chars</td>
  <td align="center">&ndash;</td>
</table>

