#!/usr/bin/env bash

wget -O definitions.zip http://hl7.org/fhir/R4/definitions.json.zip
unzip definitions.zip profiles-types.json valuesets.json -d fhir
rm definitions.zip
wget -O fhir/bundle.json http://hl7.org/fhir/R4/bundle.profile.json
wget -O fhir/codesystem.json http://hl7.org/fhir/R4/codesystem.profile.json
wget -O fhir/structuredefinition.json http://hl7.org/fhir/R4/structuredefinition.profile.json
wget -O fhir/valueset.json http://hl7.org/fhir/R4/valueset.profile.json

# Apply local overrides on top of the official spec (e.g. codes we need that
# aren't in the upstream CodeSystem). Resources are indexed by canonical URL
# during generation and later files win, so copying these in last with a
# "zz-" prefix (sorting after valuesets.json) makes them take precedence.
for f in overrides/*.json; do
  cp "$f" "fhir/zz-$(basename "$f")"
done

go generate ./fhir
