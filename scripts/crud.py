#!/usr/bin/python3
import os
import jinja2

def render_template(template_path, output_path, context):
    with open(template_path, 'r') as file:
        template_content = file.read()

    template = jinja2.Environment(loader=jinja2.FileSystemLoader('/')).from_string(template_content)
    rendered_content = template.render(context)
    with open(output_path, 'w') as file:
        file.write(rendered_content)

def setup_crud(template_dir, output_dir, context):
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    for root, _, files in os.walk(template_dir):
        relative_path = os.path.relpath(root, template_dir)
        output_path = os.path.join(output_dir, relative_path)

        if not os.path.exists(output_path):
            os.makedirs(output_path)

        for file_name in files:
            if file_name.endswith(".j2"):
                template_path = os.path.join(root, file_name)
                output_file_name = file_name[:-3] + ".go"
                output_file_path = os.path.join(output_path, output_file_name)
                render_template(template_path, output_file_path, context)

def main():
    while True:
        crud_name = input("What name for the CRUD ? First letter should be uppercase.\n")
        if not crud_name or not crud_name[0].isupper():
            print(f"'{crud_name}' is not a valid name for the CRUD.")
        else:
            break

    crud_context = {
        "name": crud_name,
        "name_lower": crud_name.lower()
    }
    template_dir = "templates"
    output_dir = "internal"

    # Setup CRUD files
    setup_crud(f"{template_dir}/crud/", f"{output_dir}/{crud_name.lower()}", crud_context)

    # Setup models for the CRUD
    render_template(f"{template_dir}/models/model.j2", f"{output_dir}/models/{crud_name.lower()}.go", crud_context)


if __name__ == "__main__":
    main()
