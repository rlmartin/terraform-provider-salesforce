import json, os, sys

###############################
# Small script for dynamically generating a swagger spec file.
# Requires Python3
#
# It uses the "component" spec files found in spec_files/components to aggregate the needed definitions, paths, and tags.
# In order to exlude a certain component, run the script with the component names you'd like to exclude as arguments.
# i.e. `python createSpecFile.py dashboard` will generate a spec file with all components except dashboards.
###############################

COMPONENTS_DIR = "./components"
INFO = "info"
WARN = "warn"

def log(level, log):
    print(f"[{level.upper()}]:", log)

def buildBaseJson():
    base = {}
    base["host"]: "https://vestahealthcare--uat.sandbox.my.salesforce.com"
    base["basePath"] = "/services/data/v57.0/tooling/sobjects"
    base["definitions"] = {}
    base["externalDocs"] = {
        "description": "salesforce Support Docs",
        "url": "https://www.salesforce.com/support/rest-api-developers-guide/"
    }
    base["info"] = {
      "description": "salesforce is a SaaS-based performance monitoring platform that provides full visibility into complex, hybrid infrastructures, offering granular performance monitoring and actionable data and insights. salesforce_sdk enables you to manage your Salesforce account programmatically.",
      "license": {
        "name": "Apache 2.0",
        "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
      },
      "title": "salesforce REST API",
      "version": "2.0.0"
    }
    base["paths"] = {}
    base["schemes"] = ["https"]
    base["securityDefinitions"] = {
      "LMv1": {
            "in": "header",
            "name": "Authorization",
            "type": "apiKey"
        }
    }
    base["swagger"] = "2.0"
    base["tags"] = []
    return base

def mergeSpecs(base, j):
    mergeDict(base["definitions"], j["definitions"], "definitions")
    mergeDict(base["paths"], j["paths"], "paths")
    mergeList(base["tags"], j["tags"], "tags")

def mergeDict(base, j, key):
    for key in j:
        if key in base:
            if base[key] != j[key]:
                log(WARN, f"found overlapping keys with different values. key: {key}")
            else:
                log(INFO, f"found overlapping keys with the same values. key:{key}")
            continue
        base[key] = j[key]

def mergeList(base, j, key):
    for val in j:
        if val in base:
            log(INFO, f"found duplicate value in list. key: {key}, val: {val}")
            continue
        base.append(val)

def main(argv):
    base = buildBaseJson()
    for filename in os.listdir(COMPONENTS_DIR):
        if filename.split(".")[0] in argv:
            continue
        with open(os.path.join(COMPONENTS_DIR, filename), 'r') as f:
            log(INFO, "merging "+filename)
            j = json.load(f)
            mergeSpecs(base, j)
    output = open("current.json", "w")
    output.write(json.dumps(base, sort_keys=True, indent=10))
    output.close()

main(sys.argv[1:])