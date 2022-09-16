# adf_trigger_cli
Azure ADF Trigger Cli

ADF Trigger CLI  is utility for run ADF with specific pipeline.


```
Build from source code

go build 
```

```
Pipeline's parameters

Exmaple

{
  "params":[
    {
        "key":"data_file",
        "value":"data.csv",
        "type" : "string"
    }
  ]
}

```

```
Example 

./adf_trigger_cli run \
--subscription_id=xxxx \
--resource_group=xxxx \
--factory_name=xxxx \
--pipeline_name=xxxx
--parameter_file=config.json

```
