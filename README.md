# how to separate envs

* There will be main branch containing all the resource blocks
  * There will not be any provider blocks in main branch so we won't be able to terraform apply from main
  * There will be separate branches for each and every envs.
  * All the separate branches of envs will have different providers.tf and vars.tfvars file
  * run pipelines of each branches for changes in different envs.


  ## Regarding use of kubernetes providers.
  * its hard to maintain k8s changes from repo as there will be rapid changes here.
  * If we plan to use this in Disaster recovery scenario then its okay to keep it here.
    * If not we can use something like flux or argocd to continuously sync changes from this repo to cluster.
  