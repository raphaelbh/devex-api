# app structure

## templates
- name
- repository_app
- building
   {
      on: { event: GIT_PUSH, data: { branch: develop } },
      execute: { 
         type: PIPELINE, 
         data: { 
            pipeline_id: '1', 
            input: {
         service_name: 'devex-api',
                  service_repository: 'https://github.com/raphaelbh/devex-api'
            }  
         }
      }
   }

## services
- name
- repository_app
- template_id
- env_variables (para cada ambiente cadastrado)
   { environment_id: {}, environment_id: {} }

## service_variables
- application_id
- type: VARIABLE | SECRET
- key
- value

## service_events
- application_id
- type: BUILD | DEPLOY
- trigged_at
- trigged_by
- data: {}

## environments
- name
- deployment
{
   on: { event: PIPELINE_EXECUTION_COMPLETED, data: { pipeline_id: 1 } },
   execute: { 
      type: PIPELINE, 
      data: { 
         pipeline_id: 2, 
         input: { 
            service_name: vira do evento
            tag: vira do evento
         }
      }
   }
}

----------

## pipelines 
- name
- version
- definition
   {
      "image": "alpine",
      "commands": [
         "echo go service building..."
      ]
   }

## pipeline_executions
- pipeline_id
- status [ WAITING_EXECUTION | IN_PROGRESS | CANCELED | FAIL | COMPLETED ]
- input
- logs
- elapsed_time
- starting_at
- finishing_at
- trigged_at
- trigged_by
- trigger_type [ MANUAL | PIPELINE | GIT | API ]

## pipeline_variables
- pipeline_id
- type: VARIABLE | SECRET
- key
- value