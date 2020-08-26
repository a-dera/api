1. gcloud endpoints services deploy api-description.yaml \
  --project business-282516
2. Copy the CONFIG_ID from the output
3. Run: 
    ./gcloud_build_image -s gateway-ksdxotlq6a-uc.a.run.app \
        -c 2020-08-26r0 -p business-282516
4. Copy the new image url and run: 
    gcloud run deploy gateway \
      --image="NEW_IMAGE_URL" \
      --allow-unauthenticated \
      --platform managed \
      --project=business-282516
      
      # Example:
      gcloud run deploy gateway \
        --image="gcr.io/business-282516/endpoints-runtime-serverless:2.15.0-gateway-ksdxotlq6a-uc.a.run.app-2020-08-14r2" \
        --allow-unauthenticated \
        --platform managed \
        --project=business-282516