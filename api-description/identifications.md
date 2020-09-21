
```
curl env/v1/persons/:person_id/identifications
-X POST
  -H "Content-Type: multipart/form-data"
  -F file=@localfilename
```

Example response :

```
{
  "id": "9dfe2f4edaa67138be0c0c1cd3a7d849cidt",
  "reference": null,
  "status": "created",
  "completed_at": null,
  "data": {
        "file": "https://url.to/file.pdf",
        "document_category": "others",
        "document_type": "government_id",
        "metadata": {},
        "status": "pending",
        "note": null,
        "created": 1516281408895,
        "updated": 1528454842365
    },
  "proof_of_address_type": null, 
  "proof_of_address_issued_at": null,
  "language": "EN"
}
```

Autorized Document Types :

| Value | Descritpion |
|-------|-------------|
| government_id | Government ID |
| passport | Passport |
| drivers_license | Drivers License |
