import os

def main():
    base_dir = 'internal/tools'
    for root, dirs, files in os.walk(base_dir):
        for file in files:
            if file.endswith('.go'):
                path = os.path.join(root, file)
                with open(path, 'r', encoding='utf-8') as f:
                    text = f.read()
                
                text = text.replace('"github.com/voocel/agentcore/schema"', '')
                text = text.replace('"github.com/voocel/agentcore/tools"', '')
                
                lines = text.split('\n')
                in_schema = False
                for i, line in enumerate(lines):
                    if ') Schema() map[string]any' in line:
                        in_schema = True
                    if in_schema:
                        lines[i] = '// ' + line
                        if line.startswith('}'):
                            in_schema = False
                
                with open(path, 'w', encoding='utf-8') as f:
                    f.write('\n'.join(lines))

if __name__ == '__main__':
    main()
