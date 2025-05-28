from tinytag import TinyTag
import os

PODCAST_DIR = r'/Users/samdevlin/Library/Group Containers/243LU875E5.groups.com.apple.podcasts/Library/Cache'

def print_tag_data(path):
    tag: TinyTag = TinyTag.get(path)
    metadata: dict = tag.as_dict()
    print(f"{tag.filename} details:\n")
    print(f"Podcast -> \t {tag.album}")
    print(f"Episode -> \t {tag.title}")
    print(f"Genre -> \t {tag.genre}")
    print("\n\n")


with os.scandir(PODCAST_DIR) as it:
    for entry in it:
        if entry.name.endswith('.mp3') and entry.is_file():
            print_tag_data(entry.path)