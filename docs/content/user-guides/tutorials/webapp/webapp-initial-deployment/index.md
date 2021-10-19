---
type: docs
title: "Deploy the website tutorial frontend"
linkTitle: "Deploy frontend"
description: "Deploy the website tutorial frontend in a container"
weight: 2000
slug: "frontend"
---

## Define a Radius app as a .bicep file

Radius uses the [Bicep language](https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/bicep-overview) as its file-format and structure. In this tutorial you will define an app named `webapp` that will contain the container and database resources, all described in Bicep.

Create a new file named `template.bicep` and paste the following:

{{< rad file="snippets/empty-app.bicep" embed=true >}}

## Add a container component

Next you'll add a `todoapp` resource for the website's frontend.

Radius captures the relationships and intentions behind an application, which simplifies deployment and management. The single `todoapp` resource in your template.bicep file will contain everything needed for the website frontend to run.

Your `todoapp`, which is a ContainerComponent resource, will specify:
- **container image:** `radiusteam/tutorial-todoapp`, a Docker image the container will run. This is where your website's front end code lives.
- **bindings:** `http`, a Radius binding that adds the ability to listen for HTTP traffic (on port 3000 here).[]

Update your template.bicep file to match the full application definition:

{{< rad file="snippets/container-app.bicep" embed=true >}}

## Deploy the application

Now you are ready to deploy the application for the first time.

1. Make sure you have an [Radius environment initialzed]({{< ref create-environment >}}).
   - For Azure environments run `az login` to make sure your token is refreshed.
   - For Kubernetes environments make sure your `kubectl`  context is set to your cluster.

2. Deploy to your Radius environment via the rad CLI:

   ```sh
   rad deploy template.bicep
   ```

   This will deploy the application into your environment and launch the container resource for the frontend website.

3. Confirm that your Radius application was deployed:

   ```sh
   rad resource list --application webapp
   ```

   You should see your `todoapp` resource. Example output:
   ```
   RESOURCE   TYPE
   todoapp    ContainerComponent
   ```

4. To test your `webapp` application, open a local tunnel to your application:

   ```sh
   rad resource expose ContainerComponent todoapp --application webapp --port 3000
   ```

   {{% alert title="💡 rad expose" color="primary" %}}
   The [`rad resource expose`]({{< ref rad_resource_expose >}}) command requires the resource type, resource name, application name flag, and port flag. If you changed any of these names when deploying, update your command to match.
   {{% /alert %}}

5. Visit the URL [http://localhost:3000](http://localhost:3000) in your browser. For now you should see a page like:

   <img src="todoapp-nodb.png" width="400" alt="screenshot of the todo application with no database">

   If the page you see matches the screenshot, that means the container is running as expected.

   You can play around with the application's features features:
   - Add a todo item
   - Mark a todo item as complete
   - Delete a todo item

6. When you're done testing press CTRL+C to terminate the port-forward.

<br>{{< button text="Next: Add a database to the app" page="webapp-add-database.md" >}}