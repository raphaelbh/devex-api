-- create environment
INSERT INTO public.environments
(id, created_at, updated_at, deleted_at, name, deployment)
VALUES(
    '2oOpIzeVObUdZiRNUkNvV2JYAxu', 
    CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, 
    NULL, 
    'Production', 
    NULL
);

-- create template
INSERT INTO public.templates
(id, created_at, updated_at, deleted_at, name, repository_app, building)
VALUES(
    '2oOpJ4UcEBX2KLKQxZkX2RFkrl2', 
    CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, 
    NULL, 
    'Template - Go Service', 
    'https://github.com/raphaelbh/go-api',
    '{
        "on": { "event": "GIT_PUSH", "data": { "branch": "develop" } },
        "execute": { 
            "type": "PIPELINE", 
            "data": { 
                "pipeline_id": "2oXqrrDrHElp7fxOcaK40jrN2Kz", 
                "input": {
                    "service_name": "devex-api",
                    "service_repository": "https://github.com/raphaelbh/devex-api"
                }  
            }
        }
    }'::jsonb
);

-- create service
INSERT INTO public.services
(id, created_at, updated_at, deleted_at, name, repository_app, env_variables, template_id)
VALUES(
    '2oXqruB2vbXih69117y1OtT0TAS', 
    CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, 
    NULL, 
    'devex-api', 
    'https://github.com/raphaelbh/devex-api', 
    '{}',
    '2oOpJ4UcEBX2KLKQxZkX2RFkrl2'
);

-- create pipeline
INSERT INTO public.pipelines
(id, created_at, updated_at, deleted_at, name, definition)
VALUES(
    '2oXqrrDrHElp7fxOcaK40jrN2Kz', 
    CURRENT_TIMESTAMP, 
    CURRENT_TIMESTAMP, 
    NULL, 
    'Build - Go Service',
    '{
        "image": "alpine",
        "commands": [
            "echo go service building..."
        ]
    }'::jsonb
);


